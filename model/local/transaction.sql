CREATE TABLE transaction (
  id              BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  send_node_id    BIGINT UNSIGNED NOT NULL,
  recieve_node_id BIGINT UNSIGNED NOT NULL,
  amount          DOUBLE          NOT NULL,
  created_at      DATETIME        NOT NULL
);