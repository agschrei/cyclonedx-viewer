{{template "base" .}}

{{define "content"}}
    <div class="container">
        <div class="row my-4">
            <div class="col">
                <h1>Welcome to the CycloneDx Viewer!</h1>
                <p>This web application allows you to browse the contents of a CycloneDx BOM</p>
            </div>
        </div>
        <div class="accordion" id="bomOverview">
            <div class="accordion-item">
                <h2 class="accordion-header accent" id="metadataHeading">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#bomMetadata" aria-expanded="true" aria-controls="bomMetadata">
                        BOM Metadata
                    </button>
                </h2>
                <div id="bomMetadata" class="accordion-collapse collapse show" aria-labelledby="metadataHeading" data-bs-parent="#bomOverview">
                    <div class="accordion-body">
                        <strong>This is the first item's accordion body.</strong> It is shown by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions. You can modify any of this with custom CSS or overriding our default variables. It's also worth noting that just about any HTML can go within the <code>.accordion-body</code>, though the transition does limit overflow.
                    </div>
                </div>
            </div>

        </div>
    </div>
{{end}}