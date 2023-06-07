CREATE TABLE IF NOT EXISTS hydra_oauth2_device_code 
(
    signature          VARCHAR(255) NOT NULL PRIMARY KEY,
    request_id         VARCHAR(255)  NOT NULL,
    requested_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    client_id          VARCHAR(255) NOT NULL,
    scope              TEXT         NOT NULL,
    granted_scope      TEXT         NOT NULL,
    form_data          TEXT         NOT NULL,
    session_data       TEXT         NOT NULL,
    subject            VARCHAR(255) NOT NULL DEFAULT '',
    active             BOOL         NOT NULL DEFAULT true,
    requested_audience TEXT         NULL DEFAULT '',
    granted_audience   TEXT         NULL DEFAULT '',
    challenge_id       VARCHAR(40)  NULL REFERENCES hydra_oauth2_flow (consent_challenge_id) ON DELETE CASCADE,
    nid                VARCHAR(36)  NULL REFERENCES networks(id) ON DELETE CASCADE ON UPDATE RESTRICT
);
CREATE INDEX hydra_oauth2_device_code_request_id_idx ON hydra_oauth2_device_code (request_id, nid);
CREATE INDEX hydra_oauth2_device_code_client_id_idx ON hydra_oauth2_device_code (client_id, nid);
CREATE INDEX hydra_oauth2_device_code_challenge_id_idx ON hydra_oauth2_device_code (challenge_id, nid);

CREATE TABLE IF NOT EXISTS hydra_oauth2_user_code 
(
    signature          VARCHAR(255) NOT NULL PRIMARY KEY,
    request_id         VARCHAR(40)  NOT NULL,
    requested_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    client_id          VARCHAR(255) NOT NULL,
    scope              TEXT         NOT NULL,
    granted_scope      TEXT         NOT NULL,
    form_data          TEXT         NOT NULL,
    session_data       TEXT         NOT NULL,
    subject            VARCHAR(255) NOT NULL DEFAULT '',
    active             BOOL         NOT NULL DEFAULT true,
    requested_audience TEXT         NULL DEFAULT '',
    granted_audience   TEXT         NULL DEFAULT '',
    challenge_id       VARCHAR(40)  NULL REFERENCES hydra_oauth2_flow (consent_challenge_id) ON DELETE CASCADE,
    nid                VARCHAR(36)  NULL REFERENCES networks(id) ON DELETE CASCADE ON UPDATE RESTRICT
);
CREATE INDEX hydra_oauth2_user_code_request_id_idx ON hydra_oauth2_user_code (request_id, nid);
CREATE INDEX hydra_oauth2_user_code_client_id_idx ON hydra_oauth2_user_code (client_id, nid);
CREATE INDEX hydra_oauth2_user_code_challenge_id_idx ON hydra_oauth2_user_code (challenge_id, nid);