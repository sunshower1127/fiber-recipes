prefork 켜면 멀티 프로세스모드가 됨. 그냥 일단 코어수대로 fork하고 봄.

production이 멀티코어 환경이면 유용하겠네요 ㅇㅇ

prefork가 false면 그냥 싱글 프로세스임. 중간에 프로세스를 알아서 fork하고 그러진 않는다고 함.
