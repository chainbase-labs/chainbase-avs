FROM repository.chainbase.com/network/flink:v1.18-0.2
# Use the official Flink image as base for the final image

# Copy the built JAR from the build stage
ADD node/flink/proof/target/proof-1.0-SNAPSHOT.jar /opt/proof/proof-1.0-SNAPSHOT.jar