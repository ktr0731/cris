<app>
    <material-navbar>
        <div class="logo"><a href="#">CRIS</a></div>
    </material-navbar>

    <div id="container">
        <p><span>ETH : </span><span>{ address }</span></p>

        <upload></upload>

        <status></status>
    </div>

    <script>
        import showStatus from '../js/actions/status';

        const self = this;
        self.opts = opts;

        opts.store.on('get_account', res => {
            self.address = res.address;
            self.update();
        })

        self.on('mount', () => {
            showStatus(opts.store, opts.ethClient);
        })

    </script>

    <style>
        material-navbar {
            padding: 0 350px;
        }
    </style>
</app>
