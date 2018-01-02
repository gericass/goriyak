CREATE TABLE transaction (
  id              BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  transaction_id  VARCHAR(65) UNIQUE,
  send_node_id    VARCHAR(65) NOT NULL,
  recieve_node_id VARCHAR(65) NOT NULL,
  amount          DOUBLE      NOT NULL,
  created_at      DATETIME    NOT NULL
);