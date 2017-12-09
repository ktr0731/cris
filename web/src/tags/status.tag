<status>
    <p>Uploaded files:</p>
    <table>
        <tr each={ files }>
            <td class="name">
                <button data-clipboard-text={"http://localhost:3000/?url=" + Base64.encodeURI(token + "." + hash + "." + localStorage['privkey'])}></button>
                { name }
            </td>
            <td class="hash">{ txHash.slice(0, 31) }</td>
        </tr>
    </table>

    <script>
        import { Base64 } from 'js-base64';
        import Clipboard from 'clipboard';
        import 'riot';

        const self = this;

        new Clipboard('button');

        self.on('mount', () => {
            self.files = [];

            self.parent.opts.store.on('set_uploaded_files', files => {
                self.files = files;
                self.update();
            })

            self.parent.opts.store.on('add_uploaded_file', file => {
                self.update();
            });
        });
    </script>

    <style>
        table {
            width: 100%;
        }

        tr {
            width: 100%;
        }

        td {
            display: inline-block;
            text-overflow: ellipsis;
            overflow: hidden;
            width: 48%;
        }
    </style>

</status>