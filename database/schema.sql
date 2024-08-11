CREATE TABLE `ffmpeg-video-database`.video
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'video 테이블의 기본 키',
    filename   VARCHAR(255) NOT NULL COMMENT '업로드된 비디오의 원본 파일 이름',
    file_path  VARCHAR(255) NOT NULL COMMENT '업로드된 비디오가 저장된 파일 경로',
    duration   BIGINT NOT NULL COMMENT '업로드된 비디오의 길이(초 단위)',
    status     ENUM('uploaded', 'processed', 'failed', 'finished') DEFAULT 'uploaded' NULL COMMENT '비디오의 처리 상태',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL COMMENT '비디오 레코드가 생성된 시간',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '비디오 레코드가 마지막으로 수정된 시간'
) COMMENT='업로드된 비디오의 세부 정보를 저장하는 테이블';

CREATE TABLE `ffmpeg-video-database`.video_job
(
    id               BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'video_job 테이블의 기본 키',
    video_id         BIGINT NOT NULL COMMENT 'video 테이블의 비디오를 참조하는 외래 키',
    job_type         ENUM('trim', 'concat') NOT NULL COMMENT '비디오에 수행할 작업 유형(트림 또는 이어붙이기)',
    parameters       LONGTEXT COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`parameters`)) COMMENT '작업에 필요한 매개변수를 저장하는 JSON 문자열',
    status           ENUM('pending', 'in_progress', 'completed', 'failed') DEFAULT 'pending' NULL COMMENT '비디오 작업의 현재 상태',
    created_at       DATETIME DEFAULT CURRENT_TIMESTAMP NULL COMMENT '비디오 작업이 생성된 시간',
    updated_at       DATETIME DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '비디오 작업이 마지막으로 수정된 시간'
) COMMENT='비디오에 수행될 작업(트림 또는 이어붙이기)의 세부 정보를 저장하는 테이블';

CREATE TABLE `ffmpeg-video-database`.final_video
(
    id                BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'final_video 테이블의 기본 키',
    original_video_id BIGINT NOT NULL COMMENT 'video 테이블의 원본 비디오를 참조하는 외래 키',
    filename          VARCHAR(255) NOT NULL COMMENT '최종 처리된 비디오의 파일 이름',
    file_path         VARCHAR(255) NOT NULL COMMENT '최종 처리된 비디오가 저장된 파일 경로',
    duration          BIGINT NOT NULL COMMENT '최종 처리된 비디오의 길이(초 단위)',
    status            ENUM('processed', 'failed') DEFAULT 'processed' NULL COMMENT '최종 비디오 처리 상태',
    created_at        DATETIME DEFAULT CURRENT_TIMESTAMP NULL COMMENT '최종 비디오 레코드가 생성된 시간',
    updated_at        DATETIME DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '최종 비디오 레코드가 마지막으로 수정된 시간'
) COMMENT='트리밍 또는 이어붙이기 작업 후 최종 처리된 비디오의 세부 정보를 저장하는 테이블';