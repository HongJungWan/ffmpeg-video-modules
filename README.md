# ffmpeg-video-modules

### FFmpeg 및 FFprobe 설치 가이드 (Windows 10)

```
# 📌 FFmpeg 및 FFprobe 설치
  - [FFmpeg 다운로드 페이지](https://ffmpeg.org/download.html)로 이동
  - Windows 빌드 선택 (권장: [gyan.dev](https://www.gyan.dev/ffmpeg/builds/))
  - "Release builds"에서 `ffmpeg-release-essentials.zip` 다운로드
  - ZIP 파일을 `C:\ffmpeg`에 압축 해제


# 📌 PATH 환경 변수 추가
  - `C:\ffmpeg\bin` 경로 복사
  - 내 PC → 속성 → 고급 시스템 설정 → 환경 변수 클릭
  - 시스템 변수에서 `Path` 편집 → 새 경로에 `C:\ffmpeg\bin` 추가
  - 모든 창 닫기


# 📌 설치 확인
  - Win + R → "cmd" 입력 후 실행
  - `ffmpeg -version` 및 `ffprobe -version` 입력해 설치 확인
```

* 위 과정을 진행해야 Local 환경에서 ffmpeg 라이브러리를 사용할 수 있습니다.

<br><br>

### 실행

```
# 📌 GoLand IDE 기준
Program arguments : -c app/deploy/dev/config.toml
```

<br><br>

### 테스트 코드 실행 시키기

```
# 📌 Git Bash 기준 명령어


# 📌 특정 디렉토리의 모든 테스트 코드 실행
go test ./app/...


# 📌 커버리지 프로파일 생성 후, HTML 보고서 생성
go test -v -coverprofile=coverage.out ./app/...
go tool cover -html=coverage.out
```

<br><br>

### 헬스 체크 API 테스트

```git bash
$ curl --http1.1 http://localhost:3031/api/health
```

<br><br>

### 컨트롤러 메서드 명명 규칙

| Controller Method | HTTP Method | Description                           |
|-------------------|-------------|---------------------------------------|
| `Index`           | GET         | 리소스의 목록을 표시합니다.              |
| `Store`           | POST        | 새로 생성된 리소스를 저장소에 저장합니다. |
| `Show`            | GET         | 지정된 리소스를 표시합니다.              |
| `Update`          | PUT/PATCH   | 지정된 리소스를 저장소에서 업데이트합니다. |
| `Destroy`         | DELETE      | 저장소에서 지정된 리소스를 제거합니다.    |

<br>

### 리포지토리 메서드 명명 규칙

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

### Git 커밋 메시지 규칙

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
