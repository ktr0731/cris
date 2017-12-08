<status>
    <p>Uploaded files:</p>
    <table>
        <tr each={ files }>
            <td class="name">{ name }</td><td class="hash">{ "0x" + hash.slice(0, 31) }</td>
        </tr>
    </table>

    <script>
        import 'riot';

        const self = this;

        self.on('mount', () => {
            self.files = [];
            self.parent.opts.store.on('add_uploaded_file', file => {
                self.files.push(file);
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