--Drop associative tables first since there are foreign constraints:
DROP TABLE IF EXISTS "OrderProduct";
DROP TABLE IF EXISTS "UserToUser";
--Drop the rest of the tables:
DROP TABLE IF EXISTS "Order";
DROP TABLE IF EXISTS "Product";
DROP TABLE IF EXISTS "User";