# Batch vs Stream Processing

## Batch

Examples of processing: payroll, billing, orders from customers.

* Send a batch at a time to be processed
* Can be scheduled
* Not real-time
* Ideal for very large processing workloads
* More cost-effective

Example: Apache Hadoop

## Streaming

Examples of processing: log monitoring, customer behavior, fraud detection.

* Producers send data to the stream
* Consumers pull data from the stream
* More expensive than batch processing

Example: Apache Spark
