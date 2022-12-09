# Core Data Concepts

## Data-Related Azure Services

* Azure Storage Accounts
* Azure Blob Storage
* Azure Tables
* Azure Files
* Azure Storage Explorer
* Azure Synapse Analytics
* CosmoDB
* Azure Data Lake Store
* Azure Data Analytics
* Azure Data Box

## Roles

* Database Administrator
* Data Engineer
* Data Analyst

## Data Overview

* __Data__ - units of information
* __Data Documents__ - types of abstract groupings and data
* __Data Sets__ - unstructured logical grouping of data
* __Data Structures__ - structured data
* __Data Types__ - how single units of data are intended to be used

# Data Documents

A data document defines the collective form in which data exists

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

### Data Structures

>Data that is organized in a specific format for easy access and modification

* Unstructured
  * Microsoft Sharepoint
  * Azure Blob Storage
  * Azure Files
  * Azure Data Lake
* Semi-structured
  * Azure Tables
  * Azure Cosmos DB
  * Mongo DB
  * Apache Cassandra
* Structured
  * Postgres
  * MySQL
  * Azure SQL
  * Azure Synapse Analytics

__Unstructured__

>Files in a folder, with random files and folders

* Loose data
* Not optimized for search or analysis

__Semi-structured__

>JSON and XML files

* Has no schema
* Has some form of relationship
* Has limitations on search
* Has fields but...
  * Don`t have to be the same in every entity
  * Are only defined when needed depending on the entity

File formats

* JSON (JavaScript Object Notation)
  * Easy for humans to read and write
  * Easy parsing and generation
* ORC (Apache Optimized Row Columnar)
  * Columnar reads, predictive pushdown and lazy reads
  * Organized in stripes of data (up to 250 MB)
  * File footer, stripe footer, raw data, index data
* Parquet
  * Columnar storage
  * Efficient compression and encoding schemes
  * Record shredding and assembly algorithm
* AVRO
  * Row-based
  * Compact, fast, binary data format
  * Container file
  * RPC

__Structured__

>Relational database

* Has a schema
* Has relationships
* Easy to search for related data
* Most common is tabular data

## Datasets

*MNIST Dataset*

>A logical grouping of data

## Notebooks

*Jupyter Notebook*

>Data that is arranged in pages, designed for easy consumption

# Data Type

A single unit of data that describes how the data is intended to be used

* Numeric
  * Integer
  * Float (Decimal)
* Text Data Types
  * Character
  * String
* Composite
  * Array
  * Hash (Dictionary)
* Binary
* Boolean
* Enumeration

# Concepts

* Batch and Streaming Data
* Relational and Non Relational
* Data Modelling
* Schema and Schemaless
* Data Integrity and Data Corruption
* Normalized and Denormalized

## Batch vs Stream Processing

### Batch

Examples of processing: payroll, billing, orders from customers.

* Send a batch at a time to be processed
* Can be scheduled
* Not real-time
* Ideal for very large processing workloads
* More cost-effective

### Streaming

Examples of processing: log monitoring, customer behavior, fraud detection.

* Producers send data to the stream
* Consumers pull data from the stream
* More expensive than batch processing

>Batch data can also be transformed into streaming data and vice-versa.

## Relational Data

* __Tables__ logical grouping of rows
* __Views__ result set of a stored query on data stored in memory (temporary table)
* __Materialized Views__ result set stored query on data stored on disk
* __Indexes__ copy of the data sorted by one or multiple columns for faster reads at the cost of storage
* __Constraints__ rules applied to writes, that can ensure data integrity
* __Triggers__ a function that is triggered on a specific database event
* __Primary Key__ a column which holds the value of the primary key from another key to establish a relationship

### Relationships

* One-to-one
* One-to-many
* Many-to-many
* Many-to-many via a junction table

## Schema vs Schemaless

A schema is a formal language which describes the structure of data and can define many different data structures, serving different purposes.

It is schemaless when the cell can accept many data types, allowing developers to forgo upfront data modelling.

## Query and Querying

A query is a request for data results (reads) to perform operations like inserting, updating or deleting data (writes).

* Query
* Data Result
* Querying
* Query Language

### Row-store vs Column-store

* __Row-store__
  * Traditional relational databases
  * Good for general purpose databases
  * Suited for Online transaction processing
  * Not the best at analytics or massive amounts of data
* __Column-store__
  * Faster at aggregating values at analytics
  * Usually NoSQL-like databases
  * Good for vast amount of data
  * Suited for Online analytical processing

### Database Indexes

Improves the speed of reads from the table by storing redundant data (usually a B-Tree in disk).

## Data Integrity vs Data Corruption

* __Data integrity__ is the maintenance and assurance of data accuracy and consistency over its entire life-cycle.
* __Data corruption__ is the act or state of data not being in the intended state that will result in data loss or misinformation

## Normalized vs Denormalized

* __Normalized__ A schema designed to store non-redundant and consistent data
  * Data integrity is maintained
  * Little redundant data
  * Many tables
  * Optimized storage
* __Denormalized__ A schema that combines data so that accessing data is fast
  * Data integrity is not maintained
  * Redundant data is common
  * Fewer tables
  * Storage is less optimal

## Pivot Table

A table of statistics that summarizes (by __aggregation__) the data of a more complex extensive table. Used in data processing to draw attention to useful information and find figures and facts quickly.

## Strongly Consistent vs Eventually Consistent

* __Strongly Consistent__ means that every time you request data, you can expect consistent data to be return within `x` time, possible seconds or more
* __Eventually Consistent__ means that when you request data, you may get inconsistent data, but it will eventually be consistent

## Synchronous vs Asynchronous

* __Synchronous__ continuous stream of data that is synchronized by a timer (guarantee of time)
* __Asynchronous__ continuous stream of data separated by start and stop bits (no guarantee of time)

## Non Relational Data

Non-tabular form and optimized for different data structures

* Key/Value
* Document
* Columnar
* Graph
* A combination of these

# Data Mining

>Extraction of patterns and knowledge from large amounts of data (not the extraction of data itself)

__CRISP-DM__

1. __Business understanding__ (what does the business need)
2. __Data understanding__ (what data there is and what data is needed)
3. __Data preparation__ (how to organize data for modeling)
4. __Modeling__ (what techniques should be applied)
5. __Evaluation__ (which data model best meets the business objectives)
6. __Deployment__ (how do people access the data)

![Chart](../../../assets/crisp-dm.png)

## Data Mining Methods

* Classification
* Clustering
* Regression
* Sequential
* Association rules
  * Support
  * Confidence
  * Lift
  * Conviction
* Outer detection
* Prediction

# Data Wrangling

>Process of transforming and mapping data from one raw data form into another format with the intent of making it more appropriate and valuable for other purposes

1. Discovery
2. Structuring
3. Cleaning
4. Enriching
5. Validating
6. Publishing

# Data Modeling

[Source](https://powerbi.microsoft.com/en-us/what-is-data-modeling/)

>Define and analyze the data requirements needed to support the business process

## Data Model

>Organizes elements of data and standardizes how they relate to one another

* Analyze and define the data that the business collects and produces
* Relationships between the data
* Create visual representations of data as it’s used at the business
* Exercise the understanding and clarification of data requirements

## Types of data modeling

* Conceptual
  * How data is represented at the organization level
* Logical
  * How data is represented in software
* Physical
  * How data is physically stored

### Conceptual data modeling

A conceptual data model defines the overall structure of your business and data. It’s used for organizing business concepts, as defined by your business stakeholders and data architects. For instance, you may have customer, employee, and product data, and each of those data buckets, known as entities, has relationships with other entities. Both the entities and the entity relationships are defined in your conceptual model.

### Logical data modeling

A logical data model builds on the conceptual model with specific attributes of data within each entity and specific relationships between those attributes. For instance, Customer A buys Product B from Sales Associate C. This is your technical model of the rules and data structures as defined by data architects and business analysts, and it will help drive decisions about what physical model your data and business needs require.

### Physical data modeling

A physical data model is your specific implementation of the logical data model, and it’s created by database administrators and developers. It is developed for a specific database tool and data storage technology, and with data connectors to serve the data throughout your business systems to users as needed. This is the “thing” the other models have been leading to—the actual implementation of your data estate.

# ETL vs ELT

[Text](../../../texts/etl_vs_elt.md)

# KPI

[Text](../../../texts/kips.md)

# Data Analytics

[Text](../../../texts/data_analytics.md)
