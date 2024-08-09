CREATE TABLE video (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Filename VARCHAR(255) NOT NULL,
    FilePath VARCHAR(255) NOT NULL,
    Duration INT,
    Status ENUM('uploaded', 'processed', 'failed') DEFAULT 'uploaded',
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE final_video (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    OriginalVideoID INT, -- 원본 동영상 ID 참조
    Filename VARCHAR(255) NOT NULL,
    FilePath VARCHAR(255) NOT NULL,
    Duration INT,
    Status ENUM('processed', 'failed') DEFAULT 'processed',
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (OriginalVideoID) REFERENCES video(ID) ON DELETE CASCADE -- 원본 동영상과 관계 설정
);

CREATE TABLE video_job (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    VideoID INT,
    JobType ENUM('trim', 'concat') NOT NULL,
    Parameters JSON NOT NULL,
    Status ENUM('pending', 'in_progress', 'completed', 'failed') DEFAULT 'pending',
    ResultFilePath VARCHAR(255),
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (VideoID) REFERENCES video(ID) ON DELETE CASCADE
);
