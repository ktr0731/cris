<status>
    <p>{ address } ({ pubkey })</p>

    <script>
        import 'riot';

        const self = this;

        init(store) {
            console.log(store)

            self.store = store;

            self.store.on('get_account', res => {
                self.address = res.address;
                self.pubkey = res.pubkey;
                self.update();
            })

        }
    </script>
</status>