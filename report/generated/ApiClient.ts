/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { BaseHttpRequest } from './core/BaseHttpRequest';
import type { OpenAPIConfig } from './core/OpenAPI';
import { NodeHttpRequest } from './core/NodeHttpRequest';
import { AlertService } from './services/AlertService';
import { AlertConfigService } from './services/AlertConfigService';
import { AlertSubscriptionService } from './services/AlertSubscriptionService';
import { AutocompleteService } from './services/AutocompleteService';
import { AwareService } from './services/AwareService';
import { CollectionGroupsService } from './services/CollectionGroupsService';
import { DataloggerService } from './services/DataloggerService';
import { DistrictRollupService } from './services/DistrictRollupService';
import { DomainService } from './services/DomainService';
import { EquivalencyTableService } from './services/EquivalencyTableService';
import { EvaluationService } from './services/EvaluationService';
import { ExplorerService } from './services/ExplorerService';
import { FormulaService } from './services/FormulaService';
import { HeartbeatService } from './services/HeartbeatService';
import { HomeService } from './services/HomeService';
import { InstrumentService } from './services/InstrumentService';
import { InstrumentConstantService } from './services/InstrumentConstantService';
import { InstrumentGroupService } from './services/InstrumentGroupService';
import { InstrumentIpiService } from './services/InstrumentIpiService';
import { InstrumentNoteService } from './services/InstrumentNoteService';
import { InstrumentSaaService } from './services/InstrumentSaaService';
import { InstrumentStatusService } from './services/InstrumentStatusService';
import { MeasurementService } from './services/MeasurementService';
import { MeasurementInclinometerService } from './services/MeasurementInclinometerService';
import { MediaService } from './services/MediaService';
import { OpendcsService } from './services/OpendcsService';
import { PlotConfigService } from './services/PlotConfigService';
import { ProfileService } from './services/ProfileService';
import { ProjectService } from './services/ProjectService';
import { ProjectRoleService } from './services/ProjectRoleService';
import { ReportConfigService } from './services/ReportConfigService';
import { SearchService } from './services/SearchService';
import { SubmittalService } from './services/SubmittalService';
import { TimeseriesService } from './services/TimeseriesService';
import { UnitService } from './services/UnitService';
type HttpRequestConstructor = new (config: OpenAPIConfig) => BaseHttpRequest;
export class ApiClient {
    public readonly alert: AlertService;
    public readonly alertConfig: AlertConfigService;
    public readonly alertSubscription: AlertSubscriptionService;
    public readonly autocomplete: AutocompleteService;
    public readonly aware: AwareService;
    public readonly collectionGroups: CollectionGroupsService;
    public readonly datalogger: DataloggerService;
    public readonly districtRollup: DistrictRollupService;
    public readonly domain: DomainService;
    public readonly equivalencyTable: EquivalencyTableService;
    public readonly evaluation: EvaluationService;
    public readonly explorer: ExplorerService;
    public readonly formula: FormulaService;
    public readonly heartbeat: HeartbeatService;
    public readonly home: HomeService;
    public readonly instrument: InstrumentService;
    public readonly instrumentConstant: InstrumentConstantService;
    public readonly instrumentGroup: InstrumentGroupService;
    public readonly instrumentIpi: InstrumentIpiService;
    public readonly instrumentNote: InstrumentNoteService;
    public readonly instrumentSaa: InstrumentSaaService;
    public readonly instrumentStatus: InstrumentStatusService;
    public readonly measurement: MeasurementService;
    public readonly measurementInclinometer: MeasurementInclinometerService;
    public readonly media: MediaService;
    public readonly opendcs: OpendcsService;
    public readonly plotConfig: PlotConfigService;
    public readonly profile: ProfileService;
    public readonly project: ProjectService;
    public readonly projectRole: ProjectRoleService;
    public readonly reportConfig: ReportConfigService;
    public readonly search: SearchService;
    public readonly submittal: SubmittalService;
    public readonly timeseries: TimeseriesService;
    public readonly unit: UnitService;
    public readonly request: BaseHttpRequest;
    constructor(config?: Partial<OpenAPIConfig>, HttpRequest: HttpRequestConstructor = NodeHttpRequest) {
        this.request = new HttpRequest({
            BASE: config?.BASE ?? '',
            VERSION: config?.VERSION ?? '2.0',
            WITH_CREDENTIALS: config?.WITH_CREDENTIALS ?? false,
            CREDENTIALS: config?.CREDENTIALS ?? 'include',
            TOKEN: config?.TOKEN,
            USERNAME: config?.USERNAME,
            PASSWORD: config?.PASSWORD,
            HEADERS: config?.HEADERS,
            ENCODE_PATH: config?.ENCODE_PATH,
        });
        this.alert = new AlertService(this.request);
        this.alertConfig = new AlertConfigService(this.request);
        this.alertSubscription = new AlertSubscriptionService(this.request);
        this.autocomplete = new AutocompleteService(this.request);
        this.aware = new AwareService(this.request);
        this.collectionGroups = new CollectionGroupsService(this.request);
        this.datalogger = new DataloggerService(this.request);
        this.districtRollup = new DistrictRollupService(this.request);
        this.domain = new DomainService(this.request);
        this.equivalencyTable = new EquivalencyTableService(this.request);
        this.evaluation = new EvaluationService(this.request);
        this.explorer = new ExplorerService(this.request);
        this.formula = new FormulaService(this.request);
        this.heartbeat = new HeartbeatService(this.request);
        this.home = new HomeService(this.request);
        this.instrument = new InstrumentService(this.request);
        this.instrumentConstant = new InstrumentConstantService(this.request);
        this.instrumentGroup = new InstrumentGroupService(this.request);
        this.instrumentIpi = new InstrumentIpiService(this.request);
        this.instrumentNote = new InstrumentNoteService(this.request);
        this.instrumentSaa = new InstrumentSaaService(this.request);
        this.instrumentStatus = new InstrumentStatusService(this.request);
        this.measurement = new MeasurementService(this.request);
        this.measurementInclinometer = new MeasurementInclinometerService(this.request);
        this.media = new MediaService(this.request);
        this.opendcs = new OpendcsService(this.request);
        this.plotConfig = new PlotConfigService(this.request);
        this.profile = new ProfileService(this.request);
        this.project = new ProjectService(this.request);
        this.projectRole = new ProjectRoleService(this.request);
        this.reportConfig = new ReportConfigService(this.request);
        this.search = new SearchService(this.request);
        this.submittal = new SubmittalService(this.request);
        this.timeseries = new TimeseriesService(this.request);
        this.unit = new UnitService(this.request);
    }
}

