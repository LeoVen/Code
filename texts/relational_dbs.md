- [Relational Databases](#relational-databases)
  - [SQL](#sql)
    - [SQL Statement Types](#sql-statement-types)
  - [OLTP vs OLAP](#oltp-vs-olap)
  - [Read Replicas](#read-replicas)

# Relational Databases

## SQL

> Access and maintain data for a Relational Database Management System (RDBMS)

* ISO 9075
* Insert, update, delete, select

### SQL Statement Types

* Data Definition Language (DDL)
  * Database schema
  * `CREATE`, `ALTER`, `DROP`, `TRUNCATE`, `COMMENT`, `RENAME`
* Data Query Language (DQL)
  * Performing queries on data
  * `SELECT`, `SHOW`, `EXPLAIN PLAN`, `HELP`
* Data Manipulation Language (DML)
  * Manipulation of data
  * `INSERT`, `UPDATE`, `DELETE`, `MERGE-UPSERT`, `CALL`, `LOCK TABLE`
* Data Control Language (DCL)
  * Rights, permissions and other controls
  * `GRANT`, `REVOKE`
* Transaction Control Language (TCL)
  * Transactions within the database
  * `COMMIT`, `ROLLBACK`, `SAVEPOINT`, `SET TRANSACTION`

## OLTP vs OLAP

> Online Transaction Processing (OLTP)
> Online Analytical Processing (OLAP)

* __OLTP__: captures, stores and processes data from transactions in real time
  * Single data source
  * Short, but many transactions
  * Latency sensitive
  * Small payloads
  * __Use case__: general purpose
* __OLAP__: uses complex queries to analyze aggregated historical data from OLTP systems
  * Multiple data sources
  * Long, but fewer transactions
  * Throughput sensitive
  * Large payloads
  * __Use case__: analytics

## Read Replicas

> A copy of the database that is kept synched with the primary database and is used to improve read contention by offloading reads to a database dedicated to perform read operations
