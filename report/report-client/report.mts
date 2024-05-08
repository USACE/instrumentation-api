import type { Dash, PlotType, Data, Datum } from "plotly.js-dist-min";
import type {
  ApiClient,
  Measurement,
  PlotConfig,
  PlotConfigTimeseriesTrace,
} from "../generated";
import type { UUID } from "crypto";

interface TimeWindow {
  after: string | undefined;
  before: string | undefined;
}

export async function processReport(
  reportConfigId: UUID,
  apiClient: ApiClient,
  apiKey: string,
  onAddNewPlot: CallableFunction,
  onAddTraces: CallableFunction,
): Promise<void> {
  const rp = await apiClient.reportConfig.getReportConfigsPlotConfigs(
    reportConfigId,
    apiKey,
  );

  const contentDiv = document.getElementById("content");

  const pcs = rp.plot_configs ?? [];

  pcs.forEach(async (pc, idx) => {
    const { after, before } = parseDateRange(pc.date_range);

    const layout = { width: 800, height: 600 };

    const plotDiv = document.createElement("div");
    plotDiv.setAttribute("id", `plot-${idx}`);

    let gd = await onAddNewPlot(plotDiv, [], layout);

    const traces = pc.display?.traces ?? [];

    // traces are pre-sorted
    const tracePromises = traces.map((tr, idx) => {
      return async function () {
        const mm = await apiClient.timeseries.getTimeseriesMeasurements(
          tr.timeseries_id!,
          after,
          before,
          3000,
        );
        const trace = createTraceData(tr, mm.items!, pc);

        await onAddTraces(gd, trace, tr.trace_order ?? idx);
      };
    });

    Promise.all(tracePromises);

    contentDiv?.appendChild(plotDiv);
  });
}

function parseDateRange(dateStr: string | undefined): TimeWindow {
  if (dateStr === undefined) {
    return { after: undefined, before: undefined };
  }

  let a: number;
  let delta: number;
  let b = Date.now();
  let d = new Date(b);

  switch (String(dateStr).toLowerCase()) {
    case "lifetime":
      // arbirarity min date
      a = Date.parse("1800-01-01");
      break;
    case "5 years":
      delta = b - d.setUTCFullYear(d.getUTCFullYear() - 5);
      a = b - delta;
      break;
    case "1 year":
      delta = b - d.setUTCFullYear(d.getUTCFullYear() - 1);
      a = b - delta;
      break;
    default:
      const dateParts = String(dateStr).split(" ", 1);
      if (dateParts.length !== 2) {
        throw new Error("could not parse custom date string");
      }
      a = Date.parse(dateParts[0]!);
      b = Date.parse(dateParts[1]!);
  }

  let after: string | undefined;
  let before: string | undefined;

  if (a !== undefined) {
    after = new Date(a).toISOString();
  }
  if (b !== undefined) {
    before = new Date(b).toISOString();
  }

  return { after, before };
}

function createTraceData(
  tr: PlotConfigTimeseriesTrace,
  mm: Measurement[],
  pc: PlotConfig,
): Data {
  const filteredItems = mm.filter((m) => {
    if (pc.show_masked && pc.show_nonvalidated) return true;
    if (pc.show_masked && !m.validated) return false;
    else if (pc.show_masked && m.validated) return true;

    if (pc.show_nonvalidated && m.masked) return false;
    else if (pc.show_nonvalidated && !m.masked) return true;

    if (m.masked || !m.validated) return false;
    return true;
  });

  const x: Datum[] = new Array(filteredItems.length);
  const y: Datum[] = new Array(filteredItems.length);

  for (let i = 0; i < filteredItems.length; i++) {
    x[i] = mm[i]?.time as Datum;
    y[i] = mm[i]?.value as Datum;
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
