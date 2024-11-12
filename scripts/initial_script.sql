CREATE TABLE resources (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(50) NOT NULL,
                           display_name VARCHAR(50),
                           type VARCHAR(50) NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           deleted_at TIMESTAMP
);

INSERT INTO resources (name, display_name, type) VALUES
                                                     ('resource1', 'Resource One', 'TypeA'),
                                                     ('resource2', 'Resource Two', 'TypeB'),
                                                     ('resource3', 'Resource Three', 'TypeA'),
                                                     ('resource4', 'Resource Four', 'TypeC'),
                                                     ('resource5', 'Resource Five', 'TypeB');
