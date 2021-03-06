CREATE TABLE `data_type_test` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    numeric_bit bit comment 'numeric: bit',
    numeric_bit_m bit(10) comment 'numeric: bit(m)',
    numeric_tinyint tinyint comment 'numeric: tinyint',
    numeric_tinyint_m tinyint(3) zerofill comment 'numeric: tinyint(m)',
    numeric_bool bool comment 'numeric: bool',
    numeric_boolean boolean comment 'numeric: boolean',
    numeric_smallint smallint comment 'numeric: smallint',
    numeric_smallint_m smallint(10) unsigned comment 'numeric: smallint(m)',
    numeric_int int comment 'numeric: int',
    numeric_int_m int(10) comment 'numeric: int(m)',
    numeric_integer integer comment 'numeric: integer',
    numeric_bigint bigint comment 'numeric: bigint',
    numeric_decimal decimal comment 'numeric: decimal',
    numeric_decimal_m decimal(3) comment 'numeric: decimal(m)',
    numeric_decimal_m_d decimal(4,2) comment 'numeric: decimal(m,d)',
    numeric_dec dec comment 'numeric: dec',
    numeric_numeric numeric comment 'numeric: numeric',
    numeric_fixed fixed comment 'numeric: fixed',
    numeric_float float comment 'numeric: float',
    numeric_float_m_d float(4,2) comment 'numeric: float(m,d)',
    numeric_float_p float(4) zerofill comment 'numeric: float(p)',
    numeric_double_m_d double(4,2) zerofill comment 'numeric: double(m,d)',
    numeric_double_precision_m_n double precision(5,2) zerofill comment 'numeric:double precision(m,d)',
    numeric_real_m_d real(4,2) comment 'numeric: real(m,d)',
    str_char CHAR NOT NULL DEFAULT '' COMMENT 'str: char',
    str_char_m CHAR(32) NOT NULL DEFAULT '' COMMENT 'str: char(m)',
    str_char_national NATIONAL CHAR(32) NOT NULL DEFAULT '' COMMENT 'str: NATIONAL char(m)',
    str_char_charset CHAR(255)CHARACTER SET UTF8 COLLATE UTF8_GENERAL_CI COMMENT 'str: char(m) charset collate',
    str_varchar VARCHAR(128) COMMENT 'str: varchar(m)',
    str_binary BINARY COMMENT 'str: binary',
    str_binary_m BINARY(4) COMMENT 'str: binary(m)',
    str_varbinary VARBINARY(6) COMMENT 'str: varbinary(m)',
    str_tinyblob TINYBLOB COMMENT 'str: tinyblob',
    str_tinytext TINYTEXT COMMENT 'str: tinytext',
    str_tinytext_charset TINYTEXT CHARACTER SET UTF8 COLLATE UTF8_GENERAL_CI COMMENT 'str: tinytext charset collate',
    str_blob BLOB COMMENT 'str: blob',
    str_blob_m BLOB(255) COMMENT 'str: blob(m)',
    str_text TEXT CHARACTER SET UTF8 COLLATE UTF8_GENERAL_CI COMMENT 'str: text charset collate',
    str_text_m TEXT(4069) COMMENT 'str: text(m)',
    str_mediumblob MEDIUMBLOB COMMENT 'str: mediumblob',
    str_mediumtext MEDIUMTEXT COMMENT 'str: mediumtext',
    str_mediumtext_charset MEDIUMTEXT CHARACTER SET UTF8 COLLATE UTF8_GENERAL_CI COMMENT 'str: mediumtext charset collate',
    str_longblob LONGBLOB COMMENT 'str: longblob',
    str_longtext LONGTEXT COMMENT 'str: longtext',
    str_longtext_charset LONGTEXT CHARACTER SET UTF8 COLLATE UTF8_GENERAL_CI COMMENT 'str: longtext charset collate',
    str_enum ENUM('a', 'b', 'c') COMMENT 'str: enum',
    str_enum_charset ENUM('a', 'b', 'c')CHARACTER SET UTF8 COLLATE UTF8_GENERAL_CI COMMENT 'str: enum charset collate',
    str_set SET('x', 'y', 'z') COMMENT 'str: set',
    str_set_charset SET('x', 'y', 'z') CHARACTER SET UTF8 COLLATE UTF8_GENERAL_CI COMMENT 'str: set charset collate',
    time_date DATE NOT NULL DEFAULT '2020-06-01' COMMENT '时间：date 格式',
    time_datetime DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '时间：datetime 格式',
    time_datetime_fsp DATETIME(5) NOT NULL DEFAULT '2020-06-01 12:00:01.987' COMMENT '时间：datetime(fsp) 格式',
    time_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '时间：timestamp 格式',
    time_timestamp_fsp TIMESTAMP(3) NOT NULL DEFAULT '2020-06-01 12:00:01.987' COMMENT '时间：timestamp(fsp) 格式',
    time_time TIME NOT NULL COMMENT '时间：time 格式',
    time_time_fsp TIME(6) NOT NULL COMMENT '时间：time(fsp) 格式',
    time_year YEAR NOT NULL DEFAULT '2020' COMMENT '时间：year 格式',
    time_year_4 YEAR(4) NOT NULL DEFAULT '2004' COMMENT '时间：year(4) 格式',
    PRIMARY KEY (`id`)
)  ENGINE=INNODB DEFAULT CHARSET=UTF8;