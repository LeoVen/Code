# Data Storage in Cloud

## Datastores

*Azure Data Lake*

>Unstructured or semi-structured data housing

* Flat files
* E-mails
* Directory services
* Databases are a sub-set of datastores
* Designed to be across many machines

## Databases

*Azure SQL*

>Structured data that can be quickly accessed and searched

* More complex datastores
* Relational or non-relational databases
* Specialized language to query
* Specialized modelling strategies to optimize retrieval of data
* More fine tune control over the transformation of the data into useful data structures

## Data Warehouses

*Azure Synapse Analytics*

>Structured or semi-structured data for creating reports and analytics

* Terabytes and millions of rows of data
* Generally perform aggregations, hence column oriented
* Return queries very fast despite the large amount of data
* Infrequently accessed
* Generally reads from a Database and is read-only

## Data Mart

>A subset of a Data Warehouse

* Usually under 100 GB
* Allows different teams control over their own datasets
* Read-only
* Increases the frequency at which data can be accessed due to smaller datasets

## Data Lakes

*Azure Data Lake*

>A centralized storage repository that holds vast amounts of data in semi-structured or unstructured format

* Hoarding for data scientists
* Real-time analytics
* Machine Learning
* On-premise data

## Data Lakehouse

*Apache Delta Lake*

>Combines the best elements of data lakes and data warehouses
