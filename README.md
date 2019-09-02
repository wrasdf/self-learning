Personal Project [![CircleCI](https://circleci.com/gh/wrasdf/self-learning/tree/master.svg?style=svg)](https://circleci.com/gh/wrasdf/self-learning/tree/master)

### UI
Features:
  - Employees (CRUD)
  - Bootstrap with fancy ui support (Mobile & Desktop)

### API
Features:
    - Golang based API server
    - Connect with DB
    - Redis as a cache layer
    - Structured logic layers

### TODO Checkpoints:
ENV Supports:
  - Local
  - Sit
  - production

Features:
  - Local
    - Easy to Tests
    - Easy to Debug
    - Log supports

  - Sit
    - Test Passed
    - Easy to deploy
    - Log will be in cloudwatch
    - Metrics
    - TLS
    - Zero downtime
    - DR plan

  - Production
    - Test Passed
    - Easy to deploy
    - Log will be in cloudwatch
    - audit log ?
    - Metrics
    - TLS
    - Zero downtime
    - DR plan


Deploy Pipeline:
  - branch
    - Unit Test
    - Contract Test
    - API Test
    - docker-compose E2E test

  - Sit  
    - Unit Test
    - Contract Test
    - API Test
    - Compile FrontEnd to CloudFront by cfn
    - Compile Backend to Fargate by cfn
    - E2E test (happy pass)

  - Production
    - Unit Test
    - Contract Test
    - API Test
    - Compile FrontEnd to CloudFront by cfn
    - Compile Backend to Fargate by cfn
    - E2E test (happy pass)
