import type {
  Dash,
  PlotType,
  Datum,
  LayoutAxis,
  Layout,
  Shape,
  Config,
  PlotData,
} from "plotly.js-dist-min";
import type { UUID } from "crypto";
import type { paths, components } from "../generated";

type PlotConfigTimeseriesTrace =
  components["schemas"]["PlotConfigTimeseriesTrace"];
type Measurement = components["schemas"]["Measurement"];

const precip = "precipitation";

window.processReport = async (
  reportConfigId: UUID,
  baseUrl: string,
  apiKey: string,
): Promise<{ districtName: string }> => {
  const { newPlot } = await import("plotly.js-dist-min");
  const { default: createClient } = await import("openapi-fetch");

  const apiClient = createClient<paths>({ baseUrl });

  const { data: rp, error } = await apiClient.GET(
    "/report_configs/{report_config_id}/plot_configs",
    {
      params: {
        path: {
          report_config_id: reportConfigId,
        },
        query: {
          key: apiKey,
        },
      },
    },
  );

  if (error) {
    throw new Error(JSON.stringify(error));
  }

  const contentDiv = document.getElementById("content");

  const introDiv = document.createElement("div");
  introDiv.setAttribute("id", "intro");

  const authorDiv = document.createElement("div");
  authorDiv.innerText = `Report configuration created by ${rp.creator_username ?? "MIDAS"}`;
  authorDiv.setAttribute("id", "author");
  introDiv?.appendChild(authorDiv);

  const titleHeader = document.createElement("h1");
  titleHeader.innerText = `${rp.project_name ?? "MIDAS Project"}: ${rp.name ?? "Report"}`;
  titleHeader.setAttribute("id", "title");
  introDiv?.appendChild(titleHeader);

  const pcs = rp?.plot_configs ?? [];

  const p = document.createElement("p");
  p.innerText = rp.description ?? "";
  p.setAttribute("id", "description");
  introDiv?.appendChild(p);

  contentDiv?.appendChild(introDiv);

  const globalDateRange = rp.global_overrides?.date_range;
  const globalShowMasked = rp.global_overrides?.show_masked;
  const globalShowNonvalidated = rp.global_overrides?.show_nonvalidated;

  let dateRange = globalDateRange?.value;
  let showMasked = globalShowMasked?.value;
  let showNonvalidated = globalShowNonvalidated?.value;

  // There's an upper limit to how many points in an svg plotly can render.
  // As a workaround, we can use WebGL after we've crossed that threshold
  // The downside is that WebGL is pixel-based rather than vector-based like
  // svg, so the resulting files are larger and have lower resolution than
  // the normal scatter plots.
  let counter: number = 0;

  pcs.forEach(async (pc): Promise<void> => {
    const wrapperDiv = document.createElement("div");
    wrapperDiv.setAttribute("class", "plot-wrapper");

    if (!globalDateRange?.enabled) {
      dateRange = pc.date_range;
    }

    let { after, before } = parseDateRange(dateRange);

    const plotHeader = document.createElement("h1");
    wrapperDiv.setAttribute("class", "plot-header");
    plotHeader.innerText = pc.name ?? "Unnamed Plot";
    wrapperDiv?.appendChild(plotHeader);

    const traces = pc.display?.traces ?? [];
    let withPrecipitation = false;
    const data = await Promise.all(
      traces.map(async (tr): Promise<Partial<PlotData>> => {
        if (tr?.parameter === precip) {
          withPrecipitation = true;
        }
        let { data: mm, error } = await apiClient.GET(
          "/timeseries/{timeseries_id}/measurements",
          {
            params: {
              path: {
                timeseries_id: tr.timeseries_id!,
              },
              query: {
                after,
                before,
                threshold: 3000,
              },
            },
          },
        );

        if (error) {
          console.error(error);
        }

        if (!globalShowMasked?.enabled) {
          showMasked = pc.show_masked;
        }
        if (!globalShowNonvalidated?.enabled) {
          showNonvalidated = pc.show_nonvalidated;
        }

        const items = mm?.items ?? [];

        const filteredItems = items.filter((m) => {
          if (showMasked && showNonvalidated) return true;
          if (showMasked && !m.validated) return false;
          else if (showMasked && m.validated) return true;

          if (showNonvalidated && m.masked) return false;
          else if (showNonvalidated && !m.masked) return true;

          if (m.masked || !m.validated) return false;
          return true;
        });

        counter += filteredItems.length;

        if (counter < 10_000) {
          // https://community.plotly.com/t/webgl-plots-are-blurry/41716/3
          // This issue happens because WebGL plots are pixel-based. It may be worth
          // using them for larger datasets but the config.plotGlPixelRatio param at
          // high values causes larger files. Consider using WebGL based plots if there
          // is a very large dataset being processed and processing time/cpu usage is an issue.
          if (tr.trace_type === "scattergl") {
            tr.trace_type = "scatter";
          }
        }

        return createTraceData(tr, filteredItems);
      }),
    );

    const defaultCustomShapes: Partial<Shape>[] = [];
    const layout: Partial<Layout> = {
      width: 1000,
      height: 400,
      margin: {
        t: 0,
        pad: 0,
      },
      showlegend: true,
      legend: {
        orientation: "h",
        x: 0,
        y: -0.2,
      },
      xaxis: {
        range: [after, before],
        title: "Date",
        showline: true,
        mirror: true,
      },
      yaxis: {
        title: pc?.display?.layout?.yaxis_title ?? "Measurement",
        showline: true,
        mirror: true,
        domain: [0, withPrecipitation ? 0.66 : 1],
      },
      ...(pc?.display?.layout?.secondary_axis_title
        ? {
            yaxis2: {
              title: pc?.display?.layout?.secondary_axis_title,
              showline: true,
              side: "right" as LayoutAxis["side"],
              overlaying: "y1" as LayoutAxis["overlaying"],
              domain: [0, withPrecipitation ? 0.66 : 1],
            },
          }
        : {}),
      ...(withPrecipitation
        ? {
            yaxis3: {
              title: "Rainfall",
              autorange: "reversed",
              showline: true,
              mirror: true,
              domain: [0.66, 1],
            },
          }
        : {}),
      shapes: pc?.display?.layout?.custom_shapes?.reduce((filtered, shape) => {
        if (shape.enabled) {
          filtered.push({
            type: "line",
            x0: after,
            x1: before,
            y0: shape.data_point,
            y1: shape.data_point,
            line: {
              color: shape.color,
              width: 3,
            },
          });
        }
        return filtered;
      }, defaultCustomShapes),
    };

    const config: Partial<Config> = {
      staticPlot: true,
    };

    const graphDiv = document.createElement("div");
    graphDiv.setAttribute("class", "plot");
    wrapperDiv?.appendChild(graphDiv);
    await newPlot(graphDiv, data, layout, config);

    contentDiv?.appendChild(wrapperDiv);
  });

  return {
    districtName: rp?.district_name ?? "No District",
  };
};

function parseDateRange(dateStr: string | undefined): {
  before: string;
  after: string;
} {
  let a = new Date();
  let b = new Date();

  if (dateStr === undefined) {
    a.setUTCDate(a.getUTCDate() - 7);
    return {
      after: a.toISOString(),
      before: b.toISOString(),
    };
  }

  switch (String(dateStr).toLowerCase()) {
    case "lifetime":
      a = new Date(Date.parse("1800-01-01")); // arbirarity min date
      break;
    case "5 years":
      a.setUTCFullYear(a.getUTCFullYear() - 5);
      break;
    case "1 year":
      a.setUTCFullYear(a.getUTCFullYear() - 1);
      break;
    default:
      const dateParts = String(dateStr).trim().split(" ");
      if (dateParts.length !== 2) {
        throw new Error("could not parse custom date string");
      }
      a = new Date(Date.parse(dateParts[0]!));
      b = new Date(Date.parse(dateParts[1]!));
  }

  return { after: a.toISOString(), before: b.toISOString() };
}

function createTraceData(
  tr: PlotConfigTimeseriesTrace,
  mm: Measurement[],
): Partial<PlotData> {
  const x: Datum[] = new Array(mm.length);
  const y: Datum[] = new Array(mm.length);

  for (let i = 0; i < mm.length; i++) {
    x[i] = mm[i]?.time as Datum;
    y[i] = mm[i]?.value as Datum;
  }

  if (tr.parameter === precip) {
    tr.trace_type = "bar";
    tr.y_axis = "y3";
  }

  return {
    x: x,
    y: y,
    mode: `lines${tr.show_markers ? "+markers" : ""}`,
    line: {
      dash: tr.line_style as Dash,
      color: tr.color,
      width: tr.width,
    },
    marker: {
      size: Number(tr.width) ? Number(tr.width) * 2 + 3 : 5,
      color: tr.color,
    },
    name: tr.name,
    type: tr.trace_type as PlotType,
    yaxis: tr.y_axis,
    showlegend: true,
  };
}
