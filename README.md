# ffmpeg-video-modules

---

## Method Naming Conventions

### Controller Method Naming Rules

| Controller Method | HTTP Method | Description                           |
|-------------------|-------------|---------------------------------------|
| `Index`           | GET         | 리소스의 목록을 표시합니다.              |
| `Store`           | POST        | 새로 생성된 리소스를 저장소에 저장합니다. |
| `Show`            | GET         | 지정된 리소스를 표시합니다.              |
| `Update`          | PUT/PATCH   | 지정된 리소스를 저장소에서 업데이트합니다. |
| `Destroy`         | DELETE      | 저장소에서 지정된 리소스를 제거합니다.    |

<br>

### Repository Method Naming Rules

| Repository Method | Description                                           |
|-------------------|-------------------------------------------------------|
| `FindByXX`        | 주어진 XX로 식별된 엔티티를 반환합니다.                   |
| `FindAll`         | 모든 엔티티를 반환합니다.                               |
| `Save`            | 주어진 엔티티를 저장합니다.                             |
| `SaveByXX`        | 주어진 XX로 식별된 엔티티를 저장합니다.                   |
| `DeleteByXX`      | 주어진 XX로 식별된 엔티티를 삭제합니다.                   |
| `Count`           | 엔티티의 개수를 반환합니다.                             |
| `ExistsBy`        | 주어진 ID를 가진 엔티티가 존재하는지 여부를 나타냅니다.    |

<br><br>

## Rule & Convention

### Git commit message convention

| Tag        | Description                                         |
|------------|-----------------------------------------------------|
| `feat`     | 새로운 기능을 추가한 경우 사용합니다.                               |
| `fix`      | 버그를 수정한 경우 사용합니다.                                   |
| `refactor` | 코드 리팩토링한 경우 사용합니다.                                  |
| `style`    | 코드 형식, 정렬, 주석 등의 변경(동작에 영향을 주는 코드 변경 없음)한 경우 사용합니다. |
| `test`     | 테스트 추가, 테스트 리팩토링(제품 코드 수정 없음, 테스트 코드에 관련된 모든 변경에 해당)한 경우 사용합니다.                                             |
| `docs`     | 문서를 수정(제품 코드 수정 없음)한 경우 사용합니다.                                             |
| `chore`    | 빌드 업무 수정, 패키지 매니저 설정 등 위에 해당되지 않는 모든 변경(제품 코드 수정 없음)일 경우 사용합니다.                                             |

<br><br>
