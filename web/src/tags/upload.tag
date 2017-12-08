<upload>
    <material-spinner style="display: none;"></material-spinner>
    <div id="form">
        <input id="uploader" onchange={ change } type="file" name="content" style="visibility: hidden">
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


        change (e) {
            const self = this;

            const el = document.querySelector('material-spinner');
            const form = document.querySelector('#form');

            el.style.display = 'flex';
            form.style.display = 'none';


            upload(e.target.files[0])
                .then(() => {
                    setTimeout(() => {
                        el.style.display = 'none';
                        form.style.display = 'block';
                    }, 1000)
                })
                .catch(e => {
                    console.log(e)
                })
        }
    </script>

    <style>
        material-spinner {
            justify-content: center;
        }
    </style>
</upload>