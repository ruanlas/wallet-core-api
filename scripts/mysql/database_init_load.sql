USE wallet_core;

INSERT INTO payment_type(id, type_name)
VALUES
(1, 'Boleto'),
(2, 'Transferência'),
(3, 'Crédito');

INSERT INTO invoice_category(id, category)
VALUES
(1, 'Moradia'),
(2, 'Alimentação'),
(3, 'Transporte'),
(4, 'Educação'),
(5, 'Saúde'),
(6, 'Cuidado Pessoal e Beleza'),
(7, 'Lazer'),
(8, 'Vestuário'),
(9, 'Diversos'),
(10, 'Rateio'),
(11, 'Investimentos');

INSERT INTO gain_category(id, category)
VALUES
(1, 'Salário'),
(2, '13º Salário'),
(3, 'Férias'),
(4, 'Prestação de Serviços'),
(5, 'Premiação'),
(6, 'Dividendos'),
(7, 'Aluguéis'),
(8, 'Resgate de Investimentos'),
(9, 'Recebimento de Dívidas'),
(10, 'Rateio'),
(11, 'Outros');
