-- +migrate Up
ALTER TABLE attachments ADD COLUMN IF NOT EXISTS file_name VARCHAR(255) NOT NULL DEFAULT 'unknown';
ALTER TABLE attachments ADD COLUMN IF NOT EXISTS file_size BIGINT NOT NULL DEFAULT 0;
ALTER TABLE attachments ADD COLUMN IF NOT EXISTS mime_type VARCHAR(100);

CREATE INDEX IF NOT EXISTS idx_attachments_uploaded_by ON attachments(uploaded_by);

-- +migrate Down
ALTER TABLE attachments DROP COLUMN IF EXISTS file_name;
ALTER TABLE attachments DROP COLUMN IF EXISTS file_size;
ALTER TABLE attachments DROP COLUMN IF EXISTS mime_type;

DROP INDEX IF EXISTS idx_attachments_uploaded_by;