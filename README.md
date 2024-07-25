# wavynote backend with koyeb
golang으로 작성된 wavynote 백엔드 모듈을 koyeb을 통해 배포

## koyeb이란?
[Koyeb](https://www.koyeb.com/) is a developer-friendly serverless platform designed to let businesses easily deploy reliable and scalable applications globally. The platform has been created by Cloud Computing Veterans and is financially backed by industry leaders.

Koyeb allows you to deploy all kind of services including full web applications, APIs, and background workers.

Applications can be deployed using standard Docker containers or directly from your git repositories.

## 패키지 구조
* internal/gateway
  * http: wavynote REST API를 제공하기 위한 HTTP Server를 구현하는 패키지이며, 등록할 핸들러는 handler/ 디렉토리 하위에 추가함
    * restapi: REST API를 처리하기 위한 핸들러를 구현함
      * box: `받은노트 페이지` 관련 기능을 제공하는 패키지
      * opennote: `오픈노트 페이지` 관련 기능을 제공하는 패키지
      * profile: `프로필` 관련 기능을 제공하는 패키지
      * root: `나의노트 페이지` 관련 기능을 제공하는 패키지
      * search: `검색` 관련 기능을 제공하는 패키지
      * write: `노트` 관련 기능을 제공하는 패키지
* platform: internal에 포함된 패키지들에서 공통적으로 사용될 수 있는 기능을 제공하는 패키지들의 집합
  * dbmsadapter: DBMS의 종류에 관계없이 통일된 DBMS 관련 API를 제공하기 위한 패키지
    * postgres
  * logger: go.uber.org/zap 패키지를 활용하여 로깅 기능을 구현한 패키지
* wavynote: wavynote 모듈에 포함된 모든 패키지에서 공통적으로 사용되는 구조체를 정의하는 패키지
