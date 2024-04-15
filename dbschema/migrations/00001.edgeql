CREATE MIGRATION m15otgfz4vvi4urbfrqred2yu5wewzdyawedrgge6rxihwkvekeh6q
    ONTO initial
{
  CREATE ABSTRACT ANNOTATION default::origin;
  CREATE TYPE default::User {
      CREATE PROPERTY cosmo_id: std::str {
          CREATE ANNOTATION default::origin := 'COSMO';
          CREATE ANNOTATION std::description := 'The COSMO ID from the account';
          CREATE ANNOTATION std::title := 'COSMO ID';
      };
      CREATE PROPERTY nickname: std::str {
          CREATE ANNOTATION default::origin := 'Project';
          CREATE ANNOTATION std::description := 'The nickname chosen';
          CREATE ANNOTATION std::title := 'Nickname';
      };
      CREATE REQUIRED PROPERTY wallet_address: std::str {
          SET readonly := true;
          CREATE ANNOTATION default::origin := 'COSMO';
          CREATE ANNOTATION std::description := 'The wallet address associated with the account';
          CREATE ANNOTATION std::title := 'Wallet Address';
      };
  };
};
