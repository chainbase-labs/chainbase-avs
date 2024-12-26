package types

type NodeConfig struct {
	// used to set the logger level (true = info, false = debug)
	Production                     bool   `yaml:"production"`
	OperatorStateRetrieverAddress  string `yaml:"operator_state_retriever_address"`
	AVSRegistryCoordinatorAddress  string `yaml:"avs_registry_coordinator_address"`
	EthRpcUrl                      string `yaml:"eth_rpc_url"`
	EthWsUrl                       string `yaml:"eth_ws_url"`
	BlsPrivateKeyStorePath         string `yaml:"bls_private_key_store_path"`
	EcdsaPrivateKeyStorePath       string `yaml:"ecdsa_private_key_store_path"`
	CoordinatorServerIpPortAddress string `yaml:"coordinator_server_ip_port_address"`
	EigenMetricsIpPortAddress      string `yaml:"eigen_metrics_ip_port_address"`
	EnableMetrics                  bool   `yaml:"enable_metrics"`
	NodeApiIpPortAddress           string `yaml:"node_api_ip_port_address"`
	EnableNodeApi                  bool   `yaml:"enable_node_api"`
	NodeGrpcServerAddress          string `yaml:"node_grpc_server_address"`
	PostgresHost                   string `yaml:"postgres_host"`
	PostgresPort                   string `yaml:"postgres_port"`
	PostgresUser                   string `yaml:"postgres_user"`
	PostgresPassword               string `yaml:"postgres_password"`
	PostgresDatabase               string `yaml:"postgres_database"`
	JobManagerHost                 string `yaml:"job_manager_host"`
	JobManagerPort                 string `yaml:"job_manager_port"`
	ChainbaseRpcUrl                string `yaml:"chainbase_rpc_url"`
	CContractAddress               string `yaml:"c_contract_address"`
	StakingContractAddress         string `yaml:"staking_contract_address"`
}

type TaskIndex = uint32
