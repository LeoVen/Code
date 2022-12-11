- [Azure Data Lake](#azure-data-lake)
- [Azure Data Lake Analytics](#azure-data-lake-analytics)
  - [U-SQL](#u-sql)
- [Azure Storage Accounts](#azure-storage-accounts)
  - [Azure Blob Storage](#azure-blob-storage)
  - [Azure Files](#azure-files)
- [Azure Synapse Analytics](#azure-synapse-analytics)
  - [Synapse SQL](#synapse-sql)
  - [Synapse ELT](#synapse-elt)
- [Polybase](#polybase)
- [Power BI](#power-bi)
  - [Data Visualization](#data-visualization)
  - [Power BI Embedded](#power-bi-embedded)
  - [Power BI Interactive Reports](#power-bi-interactive-reports)
  - [Power BI Service](#power-bi-service)
  - [Reports vs Dashboards](#reports-vs-dashboards)
  - [Paginated Reports](#paginated-reports)

# Azure Data Lake

> Centralized data repository for unstructured data and semi-structured data

* Collect from various data sources
* Transform using ELT/ETL engines
* Distribute across various programs or APIs
* Publish datasets to meta catalogs for analysts
* Designed to handle petabytes of data and hundreds of gigabits of throughput
* Hierarchical namespace to efficiently access data to Azure Blob Storage

# Azure Data Lake Analytics

> On-demand analytics job service

Instead of deploying, configuring and tuning hardware, you write queries (via U-SQL) to transform data and extract valuable insights

## U-SQL

* Query and combine data from a variety of data sources
  * Azure Data Lake Storage
  * Azure Blob Storage
  * Azure SQL DB
  * Azure SQL Data Warehouse
  * SQL Server instances running in Azure VMs

# Azure Storage Accounts

> An umbrella service for various forms of managed storage

## Azure Blob Storage

> Optimized for storing massive amounts of unstructured data

* Block blobs
* Append blobs
* Page blobs

## Azure Files

> Managed file share in the cloud

* Centralized server for storage that allows a big shared drive across VMs
* __Mounting__: when the connection with the shared file system is established, a new folder is mounted on the directory tree
* Uses SMB and NFS networking protocols
* Fully managed and resilient
* Scripting and tooling to automate the management and creation, using Azure API or PowerShell

# Azure Synapse Analytics

> A data warehouse and unified analytics platform

It is a __Data Lakehouse__ for bringing data integration tools and big data analytics.

* Build ETL or ELT processes
* Ingest data from more than 95 data sources
* Integrated Apache Spark
* T-SQL for queries on data, warehouse and Spark engines
* Multi-language support: T-SQL, Python, Scala, Sparks SQL and .Net
* Integrated AI and BI
  * Azure ML Studio
  * Azure Cognito Services
  * Microsoft Power BI
* Integration with Apache Spark

![Azure Synapse Analytics Pic](../../assets/azure_synaps_analytics.png)

## Synapse SQL

A distributed version of T-SQL designed for data warehouse workloads.

* Extends T-SQL to address streaming and machine learning scenarios
* Built-in streaming to load data from cloud data sources into SQL tables
* Integrated AI with SQL by using ML models to score data using the `PREDICT` function
* __Serverless__ and __dedicated resource__ models
  * Serverless: unpredictable workloads
  * Dedicated SQL pools: predictable workloads

## Synapse ELT

> Perform ELT using Synapse SQL

![Azure Synapse Analytics ELT](../../assets/azure_synapse_analytics_elt.png)

# Polybase

> Data virtualization feature for SQL Server

* Allows SQL Server to query data with T-SQL directly from other sources without separately installing client connection software
  * SQL Server
  * Oracle
  * Teradata
  * MongoDB
  * Hadoop Clusters
  * Cosmos DB

# Power BI

[Business Intelligence](../../texts/bi.md)

> Tool for Visualization business data

* Can ingest data from many sources
* Integrate with Azure Services

Different components

* Power BI Desktop
* Power BI Mobile
* Power BI Service
* Power BI Embedded

## Data Visualization

* Bar and column
* Line
* Matrix
* Key influencer
* Treemap
* Scatter
* Bubble
* Dot plot
* Filled map

## Power BI Embedded

> PaaS analytics embedding solution that allows visuals, reports and dashboards to be embedded into an application

## Power BI Interactive Reports

> Interactive reports

* Sliders, knobs, buttons, etc
* Modify the underlying data models for a report

## Power BI Service

* Cloud-based service
* Create and interact with reports and Dashboards
* Dashboard Tiles
  * Snapshot of data pinned to the dashboard
* Dashboard
  * A single page, often called canvas

## Reports vs Dashboards

![Dash vs report](../../assets/dash_vs_report.png)

## Paginated Reports

> Reports designed to fit into a page format so they can be printed or shared
>
> The date is displayed as tables which can span multiple pages

* Report Definition Language (RDL)
  * XML representation of a SQL Server Reporting Services report definition
  * Contains data retrieval and layout information for a report
* Power BI Report Builder
  * Design pixel-perfect paginated reports
