<app>
    <material-navbar>
        <div class="logo"><a href="#">CRIS</a></div>
    </material-navbar>

    <status></status>

    <div id="container">
        <upload></upload>
    </div>

    <script>
        import Web3 from 'web3';
        import 'riot';
        const web3 = new Web3(Web3.currentProvider);
        import showStatus from '../js/actions/status';

        const self = this;
        self.opts = opts;

        self.on('mount', () => {
            console.log(self.tags)
            self.tags['status'].init(opts.store)

            showStatus(opts.store, opts.ethClient)
        })
    </script>

    <style>
        material-navbar {
            padding: 0 350px;
        }
    </style>
</app>
