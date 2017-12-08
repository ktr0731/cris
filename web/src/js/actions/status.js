export default (store, ethClient) => {
    let account = {
        address: localStorage['address'],
        pubkey: localStorage['pubkey'],
        privkey: localStorage['privkey']
    };

    if (!account.address || !account.pubkey || !account.privkey) {
        ethClient.createAccount().then(res => {
            localStorage['address'] = res.address;
            localStorage['pubkey'] = res.pubkey;
            localStorage['privkey'] = res.privkey;
            account = res;
        });
    }

    store.setAccount(account);
};
