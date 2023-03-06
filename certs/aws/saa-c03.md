- [SAA-C03](#saa-c03)
- [Main Topics](#main-topics)
- [AWS Well-Architected Framework](#aws-well-architected-framework)
- [Services](#services)
  - [IAM](#iam)
  - [CloudWatch](#cloudwatch)
  - [S3](#s3)
    - [Data Consistency](#data-consistency)
    - [Features](#features)
    - [Storage Classes](#storage-classes)
    - [Charges](#charges)
    - [Encryption](#encryption)
    - [Versioning](#versioning)
    - [Lifecycle Rules](#lifecycle-rules)
    - [Object Lock](#object-lock)
    - [Performance](#performance)
    - [Other Features](#other-features)
    - [Sharing buckets across accounts](#sharing-buckets-across-accounts)
    - [Cross region replication](#cross-region-replication)
    - [Transfer Acceleration](#transfer-acceleration)
  - [CloudFront](#cloudfront)
    - [Signed URLs and Cookies](#signed-urls-and-cookies)
  - [Snowball](#snowball)
  - [Storage Gateway](#storage-gateway)
  - [S3 Select vs. Athena](#s3-select-vs-athena)
  - [EC2](#ec2)
    - [AMI](#ami)

# SAA-C03

> AWS Solutions Architect Associate

Links

* [Certification](https://aws.amazon.com/certification/certified-solutions-architect-associate/)
* [Exam Guide](https://d1.awsstatic.com/training-and-certification/docs-sa-assoc/AWS-Certified-Solutions-Architect-Associate_Exam-Guide.pdf)
* [Sample Questions](https://d1.awsstatic.com/training-and-certification/docs-sa-assoc/AWS-Certified-Solutions-Architect-Associate_Sample-Questions.pdf)
* [Practice](https://explore.skillbuilder.aws/learn/course/external/view/elearning/13266/aws-certified-solutions-architect-associate-official-practice-question-set-saa-c03-english?saa=sec&sec=prep)

Sources to study

* [freeCodeCamp](https://www.youtube.com/watch?v=Ia-UEYYR44s&ab_channel=freeCodeCamp.org)
* [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html?did=wp_card&trk=wp_card)

Exam Domains

* 26% - Design Resilient Architectures
* 24% - Design High-Performing Architectures
* 30% - Design Secure Architectures
* 20% - Design Cost-Optimized Architectures

# Main Topics

* **Compute**
  * EC2
  * Lambda
  * Elastic Beanstalk
* **Storage**
  * S3
  * EBS
  * EFS
  * FSx
  * Storage Gateway
* **Databases**
  * RDS
  * DynamoDB
  * Redshift
* **Networking**
  * VPCs
  * Direct Connect
  * Route 53
  * API Gateway
  * AWS Global Accelerator

# AWS Well-Architected Framework

* Operational Excellence
* Security
* Reliability
* Performance Efficiency
* Cost Optimization
* Sustainability

# Services

## IAM

> To access the console you use an account and password combination. To access AWS programmatically you use a Key and Secret Key combination

* Users
* Groups
* Policies
* Roles

## CloudWatch

> Monitor AWS resources in real time

* Alarms
* Logs
* Metrics

## S3

> Simple Storage Services

* Object-based
  * Key: name of the object
  * Value: data, a sequence of bytes
  * Version ID: versioning
  * Metadata
  * Sub-resources
    * Access Control Lists
    * Torrents
* Bucket: a folder to store data
  * Names must be globally unique (across users, per region)
  * This generates a unique URL for manipulating the data

### Data Consistency

* Read after Write: immediate read is possible after write
* Eventual consistency for update (new version) and delete
* 99.9% availability
* 11x 9s for durability

### Features

* Tiered Storage
* Lifecycle Management
* Versioning
* Encryption
* MFA for delete
* ACL and Bucket Policies

### Storage Classes

[Link](https://aws.amazon.com/s3/storage-classes/)

* S3 Standard
* S3 Intelligent-Tiering
* S3 Standard-IA
* S3 One Zone-IA
* S3 Glacier Instant Retrieval
* S3 Glacier Flexible
* S3 Glacier Deep Archive
* S3 Outposts

### Charges

* Storage
* Requests
* Storage Management
* Data Transfer
* Transfer Acceleration
* Cross Region Replication

### Encryption

* In Transit
  * SSL/TLS
* At Rest
  * Server-side: SSE-S3, SSE-KMS, SSE-C
  * Client-side (you encrypt before storing)

### Versioning

* Stores all versions of an object
* Once enabled, can't be disabled, only suspended

### Lifecycle Rules

* Transition Actions (between storage classes)
* Expiration Actions (deletes objects)

### Object Lock

> Write once, read many (WORM)

Can be applied to a bucket or objects

* Governance Mode
  * Few users can be delete objects
* Compliance Mode
  * No user can delete objects, even the root user
* Retention period
* Legal Holds

### Performance

> mybucketname/folder1/subfolder1/filename.ext

Prefix: `/folder1/subfolder1/`

* Requests per prefix per second:
  * 3500 PUT/COPY/POST/DELETE
  * 5500 GET/HEAD
* Spread data across prefixes for better performance
* KMS has performance hard limits
* Multipart Upload
  * Recommended for files above 100 MB
  * Required for files above 5 GB

### Other Features

* S3 Select
  * Select data via SQL
* Glacier Select
  * SQL queries against Glacier directly

### Sharing buckets across accounts

* Bucket Policies & IAM
* ACLs & IAM
* Cross-account IAM Roles

### Cross region replication

* Destination bucket must enable versioning
* Replication starts for new versions of objects the moment you turn it on
* Permissions aren't replicated to the destination bucket
* Delete markers or deleting individual versions are not replicated

### Transfer Acceleration

> Uses CloudFront Edge Network to accelerate uploads to S3

## CloudFront

> Is a CDN (Content Delivery Network)

* **Edge Locations**
* **Origin** - the origin of all files the CDN will distribute
  * S3 bucket
  * EC2 instance
  * Elastic Load Balancer
  * Route 53
  * **OAI** - Origin Access Identity
* **Distribution** - Collection of Edge Locations and settings
  * Web Distribution
  * RTMP (media streaming)
  * **Invalidations** - Invalidate objects, directories
* **TTL** (time-to-live)
  * Cache can be cleared

###  Signed URLs and Cookies

* One Signed URL -> One File
* One Signed Cookie -> Multiple Files
* Limited lifetime
* Issues a request as the IAM user who creates the presigned URL
* Origin
  * EC2 - use CloudFront
  * S3 - Use S3 Signed URL

## Snowball

> Migrate data at petabyte-scale

* Import and export to S3

## Storage Gateway

> A service that connects on-premises software to cloud-based storage and provides seamless integration between the two environments

* File Gateway (NFS & SMB)
* Volume Gateway (iSCSI)
* Tape Gateway (VTL)

## S3 Select vs. Athena

* S3 Select is geared more towards structure data
* Athena is a geared more towards big data
* Macie - Helps identify PII (Personal Identifiable Information)

## EC2

[Link](https://aws.amazon.com/ec2/)

* [Pricing Models](https://aws.amazon.com/ec2/pricing/)
  * On demand
    * Low cost and flexibility
    * Short term, spikes or unpredictable workloads
    * Development and testing
  * Reserved
    * Steady state and usage
    * Reserved capacity
      * Standard Reserved Instances (75% off on demand)
      * Convertible Reserved Instances (up to 54% off on demand)
      * Scheduled Reserved Instances
  * Spot
    * Flexible start and end times
    * Low compute prices
  * Dedicated Hosts
    * Regulatory requirements
    * Purchased on-demand
* [Types of instances](https://aws.amazon.com/ec2/instance-types/)

### AMI

> Amazon Machine Image


