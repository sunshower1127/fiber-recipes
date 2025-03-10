데코레이터(애노테이션)도 없는데 의존성 주입? 안하는게 좋지 않을까~

# Parsley를 사용한 의존성 주입 예제

이 코드는 Go 언어의 Fiber 웹 프레임워크를 사용하면서 Parsley라는 의존성 주입(DI, Dependency Injection) 프레임워크를 적용한 예제입니다.

## 주요 개념

### 의존성 주입(Dependency Injection)

의존성 주입은 컴포넌트가 직접 의존성을 생성하는 대신, 외부에서 의존성을 제공받는 설계 패턴입니다. 이는 코드의 결합도를 낮추고 테스트 용이성을 높입니다.

### Parsley 프레임워크

Parsley는 Go 언어를 위한 의존성 주입 프레임워크로, 다음 기능을 제공합니다:

- 서비스 등록 및 관리
- 서비스 수명 주기 관리 (싱글톤, 트랜지언트 등)
- 모듈식 설계

## 프로젝트 구조

이 예제는 다음과 같은 계층 구조로 설계되었습니다:

1. **모듈(Modules)**: 서비스 구성 및 등록 담당

   - fiber_module.go: Fiber 앱 구성
   - `greeter_module.go`: 인사 서비스 등록
   - route_handlers_module.go: 라우트 핸들러 등록

2. **라우트 핸들러(Route Handlers)**: HTTP 요청 처리

   - greeter.go: `/say-hello` 엔드포인트 처리

3. **서비스(Services)**: 비즈니스 로직
   - greeter.go: 인사말 생성 로직

## 동작 방식

1. 애플리케이션 시작 시 Parsley 레지스트리에 필요한 서비스들이 등록됩니다.
2. `ConfigureFiber`는 Fiber 앱을 설정하고 라우트 핸들러를 등록합니다.
3. `/say-hello` 요청이 오면 `greeterRouteHandler`가 처리합니다.
4. 라우트 핸들러는 `Greeter` 서비스를 통해 인사말을 생성하고 사용자에게 반환합니다.

## 특징

- **느슨한 결합**: 인터페이스를 통해 구현체에 의존하지 않음
- **테스트 용이성**: 모의 객체(mock)로 쉽게 교체 가능
- **모듈화**: 기능별로 모듈 분리
- **수명 주기 관리**: 싱글톤, 트랜지언트 등 다양한 수명 주기 지원

이 예제는 Go 언어로 구조적이고 유지보수하기 쉬운 웹 애플리케이션을 만드는 방법을 보여줍니다.
