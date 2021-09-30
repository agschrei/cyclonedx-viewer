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
                <h2 class="accordion-header" id="metadataHeading">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#bomMetadata" aria-expanded="true" aria-controls="bomMetadata">
                        BOM Metadata
                    </button>
                </h2>
                <div id="bomMetadata" class="accordion-collapse collapse show" aria-labelledby="metadataHeading">
                    <div class="accordion-body">
                        <table class="table table-hover table-striped">
                            <thead>
                                <tr>
                                    <th scope="col">Attribute</th>
                                    <th scope="col">Value</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                <td>bomFormat</td>
                                <td>{{.BomFormat}}</td>
                                </tr>
                                <tr>
                                <td>specVersion</td>
                                <td>{{.SpecVersion}}</td>
                                </tr>
                                <tr>
                                <td>serialNumber</td>
                                <td>{{.SerialNumber}}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>

            <div class="accordion-item">
                <h2 class="accordion-header" id="componentStatsHeading">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#componentStatistics" aria-expanded="false" aria-controls="componentStatsHeading">
                        Component Statistics
                    </button>
                </h2>
                <div id="componentStatistics" class="accordion-collapse collapse" aria-labelledby="componentStatsHeading">
                    <div class="accordion-body">
                        <table class="table table-hover table-striped">
                            <tbody>
                                <tr>
                                <td>#Components</td>
                                <td>{{len .Components}}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
{{end}}