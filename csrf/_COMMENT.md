작동안됨
개념만 이해하고 갑시다

# CSRF (Cross-Site Request Forgery) 예제 시나리오 설명

이 코드는 CSRF(Cross-Site Request Forgery, 사이트 간 요청 위조) 취약점과 그 방어 방법을 보여주는 예제입니다. 주요 시나리오는 다음과 같습니다:

## 시나리오 개요

1. **정상 서버**(3000번 포트)와 **악의적인 서버**(3001번 포트)가 있습니다.
2. 정상 서버에는 계정 이체 기능이 있는 로그인 시스템이 있습니다.
3. 악의적인 서버는 사용자의 정상 세션을 악용하여 CSRF 공격을 시도합니다.

## 공격 시나리오

1. 사용자가 정상 서버(localhost:3000)에 "bob"/"test" 계정으로 로그인합니다.
2. 로그인 상태에서 악의적인 웹사이트(localhost:3001)를 방문합니다.
3. 악의적인 웹사이트는 자동으로 사용자 모르게 다음과 같은 공격을 수행합니다:
   - GET 요청 공격: `<img>` 태그나 링크를 통해 계정 이체 요청을 유도
   - POST 요청 공격: 숨겨진 폼을 자동으로 제출하여 이체 요청

## 취약점 시연 방법

```bash
# CSRF 보호 없이 서버 실행
go run main.go withoutCsrf
```

1. localhost:3000에 접속하고 bob/test로 로그인
2. localhost:3001 접속 시 자동으로 CSRF 공격 발생
3. 계정 잔액이 감소하는 것 확인 가능

## 보호된 버전 시연 방법

```bash
# CSRF 보호와 함께 서버 실행
go run main.go
```

1. localhost:3000에 접속하고 bob/test로 로그인
2. localhost:3001 접속 시에도 계정 잔액이 유지됨
3. CSRF 토큰 검증으로 인해 공격이 차단됨

## CSRF 보호 메커니즘

코드에서는 다음과 같은 CSRF 방어 메커니즘을 사용합니다:

```go
var csrfProtection = csrf.New(csrf.Config{
    KeyLookup:      "form:_csrf",  // 폼에서 CSRF 토큰을 찾을 위치
    CookieName:     "csrf_",      // CSRF 토큰 저장에 사용되는 쿠키 이름
    CookieSameSite: "Strict",     // 쿠키의 SameSite 속성
    Expiration:     1 * time.Hour, // 토큰 만료 시간
    KeyGenerator:   utils.UUID,    // 토큰 생성 방식
    ContextKey:     "token",       // 컨텍스트에서 토큰에 접근하는 키
})
```

이 예제는 웹 보안에서 CSRF 취약점의 위험성과 적절한 방어 메커니즘의 중요성을 보여주기 위한 교육용 시나리오입니다.
