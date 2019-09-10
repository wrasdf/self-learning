Personal Project [![CircleCI](https://circleci.com/gh/wrasdf/self-learning/tree/master.svg?style=svg)](https://circleci.com/gh/wrasdf/self-learning/tree/master)

### Deployment
Features:
  - UI -> CloudFront by CFN
  - API -> Fargate by CFN

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

### Deploy Pipeline:
  - branch
    - Unit Test
    - Contract Test ?
    - API Test
    - docker-compose E2E test

  - Sit
    - Unit Test
    - Contract Test
    - API Test
    - deploy FrontEnd to CloudFront by cfn
    - deploy Backend to Fargate | EKS by cfn
    - E2E test (happy pass)

  - Production ?

### Checkpoints:
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

  - Production ?
    - audit log ?
