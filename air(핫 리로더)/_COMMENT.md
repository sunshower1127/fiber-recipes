# air -> Hot reloader

환경변수 설정해줘야함.

.zshrc

```shell
export PATH=$PATH:$(go env GOPATH)/bin #Confirm this line in your .profile and make sure to source the .profile if you add it!!!
```

이게 GOPATH는 보통 `~/go` 입니다.
아니 gopath는 이제 안쓰기로 한거 아니였나
실행 모듈은 다운받으면 여기로 다운되네;;

---

air 실행할때 그냥

```shell
air
```

만 쳐도 잘 됨. 왜인지는 모르겠음.

---
