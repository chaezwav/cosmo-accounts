module default {
  abstract annotation origin;

  type User {
    required wallet_address: str {
        readonly := true;
        constraint exclusive;
        annotation title := 'Wallet Address';
        annotation description := "The wallet address associated with the account";
        annotation origin := 'COSMO'
    }
    
    nickname: str {
        constraint exclusive;
        annotation title := 'Nickname';
        annotation description := "The nickname chosen";
        annotation origin := 'Project'
     }
     
     cosmo_id: str {
        constraint exclusive;
        annotation title := 'COSMO ID';
        annotation description := "The COSMO ID from the account";
        annotation origin := 'COSMO'
     }
  }
};
