# CI/CD

Continuous Integration, Continuous Delivery, Continuous Deployment

## Sources

* [What is CI/CD?](https://www.redhat.com/en/topics/devops/what-is-ci-cd)
* [What is a CI/CD Pipeline?](https://www.redhat.com/en/topics/devops/what-cicd-pipeline)

## Continuous Integration

* When there are multiple developers, a CI/CD pipeline helps at making sure that the application also runs in a completely separate environment;
* Automated tests (UT, IT, etc.) make sure that new changes to the code haven't caused undesired effects or broken the application;

## Continuous Delivery

* The goal is to have a codebase that is always ready for deployment in a production environment;
* Involves test automation and code release automation, so that in the end of the process, a valid release of the application can be released to some environment

## Continuous Deployment

* Automates the release of the application to an environment;
* Can deploy the application to the cloud minutes after the code changes have been committed;
* Relies heavily on good automation design.

## Pipeline

A CI/CD Pipeline is an automated process of building, testing and deploying code. The objective is to minimize errors and maintain a consistency with software releases.

```text
   Continuous Integration     Continuous Delivery     Continuous Deployment
  +----------------------+   +-------------------+   +---------------------+
  |                      |   |  Automatically    |   |    Automatically    |
  | Build > Test > Merge | > |    release to     | > |      deploy to      |
  |                      |   |    repository     |   |      production     |
  +----------------------+   +-------------------+   +---------------------+
```
