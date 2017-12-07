<upload>
    <div>
        <input id="uploader" onchange={ change } type="file" name="content" style="display: none">
        <label for="uploader">
            <div class="uploader-wrapper">
                <p>ファイルをアップロードする</p>
            </div>
        </label>
    </div>
    <material-button class="ui">
        <div class="text">️Submit</div>
    </material-button>

    <script>
        import upload from '../js/actions/upload';

        console.log(self.opts)
        change (e) {
            console.log(e.target.files[0])

            const formData = new FormData();
            formData.append('content', e.target.files[0]);

            upload(formData)
                .then(() => {
                    console.log('DONE')
                })
        }
    </script>
</upload>