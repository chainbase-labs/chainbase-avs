package com.chainbase.manuscript;
import org.apache.flink.api.common.functions.MapFunction;
import org.apache.flink.api.common.restartstrategy.RestartStrategies;
import org.apache.flink.api.java.tuple.Tuple2;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.connector.jdbc.JdbcConnectionOptions;
import org.apache.flink.connector.jdbc.JdbcExecutionOptions;
import org.apache.flink.connector.jdbc.JdbcSink;
import org.apache.flink.streaming.api.CheckpointingMode;
import org.apache.flink.streaming.api.datastream.DataStream;
import org.apache.flink.streaming.api.environment.CheckpointConfig;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.api.functions.sink.SinkFunction;
import org.apache.flink.table.api.EnvironmentSettings;
import org.apache.flink.table.api.Table;
import org.apache.flink.table.api.bridge.java.StreamTableEnvironment;
import org.apache.flink.types.Row;

import java.math.BigInteger;
import java.security.MessageDigest;
import java.sql.Timestamp;
import java.time.Instant;
import java.time.ZoneId;

public class ProofApplication {

	public static void main(String[] args) throws Exception {
		// Check for required environment variables
		checkEnvironmentVariables();

		// Parse command line arguments
		if (args.length < 5) {
			System.err.println("Usage: FlinkJavaPowApplication <chain> <startAt> <endAt> <difficulty> <taskIndex>");
			System.exit(1);
		}

		String chain = args[0];
		long startAt = Long.parseLong(args[1]);
		long endAt = Long.parseLong(args[2]);
		int difficulty = Integer.parseInt(args[3]);
		long taskIndex = Long.parseLong(args[4]);

		// Set up execution environment
		StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();
		EnvironmentSettings settings = EnvironmentSettings.newInstance().inStreamingMode().build();
		StreamTableEnvironment tEnv = StreamTableEnvironment.create(env, settings);

		// Configure Flink environment
		configureFlink(env, tEnv);

		// Create Paimon catalog
		createPaimonCatalog(tEnv);

		// Build SQL query using parameters
		String sqlQuery = String.format("SELECT block_number, hash FROM %s.blocks WHERE block_number >= %d AND block_number <= %d", chain, startAt, endAt);

		// Execute SQL query using Table API
		Table resultTable = tEnv.sqlQuery(sqlQuery);

		// Convert result to data stream
		DataStream<Tuple2<Boolean, Row>> ds = tEnv.toRetractStream(resultTable, Row.class);

		// Apply map function
		DataStream<Row> processedStream = ds.map(new ProcessFunction(chain, difficulty, taskIndex));

		// Create PostgreSQL sink
		processedStream.addSink(createPostgreSQLSink());

		// Submit the job asynchronously
		env.executeAsync("Flink Java Proof of Work Stream Processing");
		
		System.out.println("Job submitted successfully. Exiting main method.");
	}

	private static void checkEnvironmentVariables() {
		String[] requiredEnvVars = {"OSS_ACCESS_KEY_ID", "OSS_ACCESS_KEY_SECRET"};
		for (String var : requiredEnvVars) {
			if (System.getenv(var) == null) {
				System.err.println("Error: " + var + " environment variable is not set.");
				System.exit(1);
			}
		}
	}

	private static void configureFlink(StreamExecutionEnvironment env, StreamTableEnvironment tEnv) {
		env.setParallelism(1);
		env.enableCheckpointing(60000); // 60 seconds
		env.getCheckpointConfig().setCheckpointingMode(CheckpointingMode.EXACTLY_ONCE);
		env.getCheckpointConfig().setMinPauseBetweenCheckpoints(1000);
		env.getCheckpointConfig().setCheckpointTimeout(1800000); // 30 minutes
		env.getCheckpointConfig().setMaxConcurrentCheckpoints(1);
		env.getCheckpointConfig().enableExternalizedCheckpoints(CheckpointConfig.ExternalizedCheckpointCleanup.RETAIN_ON_CANCELLATION);
		env.setRestartStrategy(RestartStrategies.fixedDelayRestart(Integer.MAX_VALUE, 10000));

		Configuration config = tEnv.getConfig().getConfiguration();
		config.setString("table.exec.sink.upsert-materialize", "NONE");
		config.setString("state.backend.type", "rocksdb");
		config.setString("state.checkpoints.dir", "file:///opt/flink/checkpoint");
		config.setString("state.savepoints.dir", "file:///opt/flink/savepoint");
		config.setString("state.backend.incremental", "true");
		config.setString("execution.checkpointing.tolerable-failed-checkpoints", "2147483647");
		config.setString("table.exec.sink.not-null-enforcer", "ERROR");
	}

	private static void createPaimonCatalog(StreamTableEnvironment tEnv) {
		tEnv.executeSql(String.format(
				"CREATE CATALOG paimon WITH (" +
						"'type' = 'paimon', " +
						"'warehouse' = 'oss://network-testnet/warehouse', " +
						"'fs.oss.endpoint' = 'network-testnet.chainbasehq.com', " +
						"'fs.oss.accessKeyId' = '%s', " +
						"'fs.oss.accessKeySecret' = '%s', " +
						"'table-default.merge-engine' = 'deduplicate', " +
						"'table-default.changelog-producer' = 'input', " +
						"'table-default.metastore.partitioned-table' = 'false', " +
						"'table-default.lookup.cache-file-retention' = '1 h', " +
						"'table-default.lookup.cache-max-memory-size' = '256 mb', " +
						"'table-default.lookup.cache-max-disk-size' = '10 gb', " +
						"'table-default.log.scan.remove-normalize' = 'true', " +
						"'table-default.changelog-producer.row-deduplicate' = 'false', " +
						"'table-default.consumer.expiration-time' = '24 h', " +
						"'table-default.streaming-read-mode' = 'file', " +
						"'table-default.orc.bloom.filter.fpp' = '0.00001', " +
						"'table-default.scan.plan-sort-partition' = 'true', " +
						"'table-default.snapshot.expire.limit' = '10000', " +
						"'table-default.snapshot.num-retained.max' = '2000'" +
						")",
				System.getenv("OSS_ACCESS_KEY_ID"),
				System.getenv("OSS_ACCESS_KEY_SECRET")
		));
		tEnv.useCatalog("paimon");
	}

	private static class ProcessFunction implements MapFunction<Tuple2<Boolean, Row>, Row> {
		private final String chain;
		private final int difficulty;
		private final long taskIndex;

		public ProcessFunction(String chain, int difficulty, long taskIndex) {
			this.chain = chain;
			this.difficulty = difficulty;
			this.taskIndex = taskIndex;
		}

		@Override
		public Row map(Tuple2<Boolean, Row> value) throws Exception {
			if (value.f0) {
				Long blockNumber = (Long) value.f1.getField(0);
				String blockHash = (String) value.f1.getField(1);
				String powResult = performProofOfWork(blockHash, difficulty);
				Timestamp insertAt = Timestamp.from(Instant.now().atZone(ZoneId.of("UTC")).toInstant());
				return Row.of(chain, blockNumber, blockHash, powResult, insertAt, difficulty, taskIndex);
			} else {
				return null;
			}
		}

		private String performProofOfWork(String inputHash, int difficulty) throws Exception {
			BigInteger target = BigInteger.ONE.shiftLeft(256 - difficulty);
			long nonce = 0;
			while (true) {
				byte[] data = (inputHash + nonce).getBytes();
				byte[] hashResult = MessageDigest.getInstance("SHA-256").digest(data);
				BigInteger hashInt = new BigInteger(1, hashResult);
				if (hashInt.compareTo(target) < 0) {
					return String.format("%064x", hashInt);
				}
				nonce++;
			}
		}
	}

	private static SinkFunction<Row> createPostgreSQLSink() {
		return JdbcSink.sink(
				"INSERT INTO pow_results (chain, block_number, block_hash, pow_result, insert_at, difficulty, task_index) " +
						"VALUES (?, ?, ?, ?, ?, ?, ?) " +
						"ON CONFLICT (chain, block_number) DO UPDATE SET " +
						"block_hash = EXCLUDED.block_hash, " +
						"pow_result = EXCLUDED.pow_result, " +
						"insert_at = EXCLUDED.insert_at, " +
						"difficulty = EXCLUDED.difficulty, " +
						"task_index = EXCLUDED.task_index",
				(statement, row) -> {
					statement.setString(1, (String) row.getField(0));
					statement.setLong(2, (Long) row.getField(1));
					statement.setString(3, (String) row.getField(2));
					statement.setString(4, (String) row.getField(3));
					statement.setTimestamp(5, (Timestamp) row.getField(4));
					statement.setInt(6, (Integer) row.getField(5));
					statement.setLong(7, (Long) row.getField(6));
				},
				JdbcExecutionOptions.builder()
						.withBatchSize(1)
						.withBatchIntervalMs(0)
						.withMaxRetries(3)
						.build(),
				new JdbcConnectionOptions.JdbcConnectionOptionsBuilder()
						.withUrl("jdbc:postgresql://postgres:5432/manuscript_node")
						.withDriverName("org.postgresql.Driver")
						.withUsername("postgres")
						.withPassword("postgres")
						.build()
		);
	}
}