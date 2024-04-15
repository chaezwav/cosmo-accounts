CREATE MIGRATION m1xdecrfxs67fn23ju4jue4zfqgwy3eaagcug56b2vbdk23fidejla
    ONTO m15otgfz4vvi4urbfrqred2yu5wewzdyawedrgge6rxihwkvekeh6q
{
  ALTER TYPE default::User {
      ALTER PROPERTY cosmo_id {
          CREATE CONSTRAINT std::exclusive;
      };
      ALTER PROPERTY nickname {
          CREATE CONSTRAINT std::exclusive;
      };
      ALTER PROPERTY wallet_address {
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
