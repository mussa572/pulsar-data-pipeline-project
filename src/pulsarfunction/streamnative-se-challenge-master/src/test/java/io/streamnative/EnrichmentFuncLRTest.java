package io.streamnative;

import io.streamnative.config.AssignmentConfig;
import io.streamnative.function.EnrichmentFunc;
import java.util.Collections;
import java.util.HashMap;
import org.apache.pulsar.common.functions.ConsumerConfig;
import org.apache.pulsar.common.functions.FunctionConfig;
import org.apache.pulsar.functions.LocalRunner;

public class EnrichmentFuncLRTest {
        public static void main(String[] args) throws Exception {
            HashMap<String, ConsumerConfig> inputSpecs = new HashMap<>();
            inputSpecs.put(
                    AssignmentConfig.INPUT_TOPIC,
                    ConsumerConfig.builder().schemaType(AssignmentConfig.SCHEMA_TYPE).build()
            );

            FunctionConfig functionConfig = FunctionConfig.builder()
                    .className(EnrichmentFunc.class.getName())
                    .inputs(Collections.singletonList(AssignmentConfig.INPUT_TOPIC))
                    .inputSpecs(inputSpecs)
                    .name("enrichment-func")
                    .runtime(FunctionConfig.Runtime.JAVA)
                    .subName("enrichment-func-subscription")
                    .cleanupSubscription(true)
                    .build();

            LocalRunner localRunner = LocalRunner.builder()
                    .brokerServiceUrl(AssignmentConfig.SERVICE_URL)
                    .functionConfig(functionConfig)
                    .build();

            localRunner.start(false);
    }
}
