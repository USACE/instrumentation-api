export function getHeaderTmpl(bgImgBase64: string) {
  return `<div style="top: 0; width: 100%; height: auto; margin: 0;">
              <img style="top: 0; max-width: 100%; max-height: 100%" src="data:image/png;base64,${bgImgBase64}" />  
          </div>`;
}

export function getFooterTmpl(svgContent: string, logoText: string) {
  return `<div style="display: inline-block; width: 100%; height: auto; margin: 0 0.7cm; font-size: 9pt;">
              <div style="position: absolute; bottom: 1pc; left: 1pc;">
                <div id="castle-logo" style="display: block; margin-bottom: 5px;">
                  ${svgContent}
                </div>
                <label for="castle-logo" style="position: absolute; bottom: 0; left: 0; overflow: hidden; white-space: nowrap;">${logoText}</label>
              </div>
              <div style="color: grey; font-style: italic; position: absolute; bottom: 1pc; right: 1pc;">
                <span class="date"></span>
                <span>&nbsp;UTC</span>
                <span style="margin-left: 25px;">Page no.&nbsp;</span>
                <span class="pageNumber"></span>
                <span>/</span>
                <span class="totalPages"></span>
              </div>
          </div>`;
}

export function getIndexHtml(orientation: "portrait" | "landscape"): string {
  return `<!doctype html>
            <html lang="en">
              <head>
                <meta charset="UTF-8" />
                <style media="print">
                  @page {
                    size: letter ${orientation};
                    margin: 4pc;
                    margin-top: 6pc;
                    font-size: 10pt;
                  }
                  body {
                    margin: 0;
                    padding: 0;
                    font-family: "Helvetica", "Arial", sans-serif;
                    color: #444444;
                    background-color: #fafafa;
                  }
                  .sheet {
                    margin: 0;
                    padding: 0;
                    overflow: hidden;
                    position: relative;
                    box-sizing: border-box;
                    page-break-after: always;
                  }
                  body.letter .sheet {
                    width: 216mm;
                    height: 280mm;
                  }
                  body.letter.landscape .sheet {
                    width: 280mm;
                    height: 216mm;
                  }
                  #content {
                    display: block;
                  }
                  #content > * {
                    display: block;
                    float: left;
                    break-inside: avoid;
                  }
                  #intro {
                    text-align: left;
                  }
                  #title {
                    font-weight: bold;
                    font-size: 24pt;
                  }
                  #author {
                    font-size: 14pt;
                    font-style: italic;
                  }
                  #description {
                    font-size: 14pt;
                  }
                  .plot-wrapper {
                    text-align: left;
                    font-weight: bold;
                    font-size: 18pt;
                  }
                  .plot-header {
                    margin-left: 30px;
                    margin-bottom: 0;
                  }
                  .plot {
                    margin-top: 0;
                  }
                </style>
                <title>MIDAS Report</title>
              </head>
              <body>
                <div class="container" id="content"></div>
              </body>
            </html>`;
}
