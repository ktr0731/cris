import 'riot';

export default class {
    constructor() {
        riot.observable(this);
    }

    saveAccount(pubkey) {}

    setAccount(account) {
        this.account = account;
        this.trigger('get_account', account);
    }
}
