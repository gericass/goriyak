CREATE TABLE transaction (
  id              BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  name            VARCHAR(65) UNIQUE, -- transaction name for confirm transaction
  send_node_id    VARCHAR(65) NOT NULL,
  receive_node_id VARCHAR(65) NOT NULL,
  amount          DOUBLE      NOT NULL,
  created_at      DATETIME    NOT NULL
);