package collector

var manualTopicMapping = map[string]string{
	"aapen-pdl-adresse-v1":                                 "pdl",
	"dittnav-brukernotifikasjon-done-v1-backup":            "personbruker",
	"nais-deployment-requests":                             "aura",
	"nais-deployment-status":                               "aura",
	"privat-tiltaksgjennomforing-godkjentAvtale":           "arbeidsgiver",
	"privat-tiltaksgjennomforing-statistikkformidling":     "arbeidsgiver",
	"privat-tiltaksgjennomforing-smsVarselResultat":        "arbeidsgiver",
	"privat-tiltaksgjennomforing-smsVarsel":                "arbeidsgiver",
	"privat-tortuga-sekvensnummerTilstand":                 "pensjonsamhandling",
	"privat-tortuga-skatteoppgjorhendelse":                 "pensjonsamhandling",
	"privat-sykepengebehandling":                           "tbd",
	"privat-syfooversikt-tilfelle-retry-v1":                "teamsykefravr",
	"privat-sf-person-v1":                                  "crm", // Team CRM?, no nais team slug,
	"privat-sf-person-pilot-v1":                            "crm", // Team CRM?, no nais team slug,
	"nks-sf-pdl-v1":                                        "crm", // Team CRM?, no nais team slug,
	"privat-person-tpsHendelse-mottatt-v1":                 "pdl",
	"privat-person-pdl-lineage-v1":                         "pdl",
	"privat-person-pdl-identhendelseSletting-v1":           "pdl",
	"privat-person-pdl-hendelsesaggregat-v2":               "pdl",
	"privat-person-pdl-fregTilbakemelding-v1":              "pdl",
	"privat-person-pdl-e500-linje-v1":                      "pdl",
	"person-privat-pdl-e500-linje-v1":                      "pdl",
	"privat-person-navMottakStatus-v1":                     "pdl",
	"privat-person-mottakEndring-v1":                       "pdl",
	"privat-person-fregTilbakemelding-mottatt-v1":          "pdl",
	"privat-pdl-sfobjects-v1":                              "pdl",
	"privat-k9-dittnav-varsel-beskjed":                     "pleiepenger",
	"privat-eessipensjon-krav-initialisering":              "eessipensjon",
	"privat-dagpenger-journalpost-mottatt-alpha":           "teamdagpenger",
	"privat-dagpenger-behov-alpha":                         "teamdagpenger",
	"postnummer":                                           "tbd",
	"pensjon-pdl-kafka-retry":                              "pensjon-saksbehandling",
	"nada-connect-status":                                  "aura",
	"nada-connect-offsets":                                 "aura",
	"nada-connect-configs":                                 "aura",
	"kafka-monitor":                                        "aura",
	"fpsoknad-mottak":                                      "teamforeldrepenger",
	"fpinfo-historikk":                                     "teamforeldrepenger",
	"dittnav-brukernotifikasjon-nyOppgave-v1-backup":       "personbruker",
	"dittnav-brukernotifikasjon-nyBeskjed-v1-backup":       "personbruker",
	"dittnav-brukernotifikasjon-cached-done-v1-backup":     "personbruker",
	"aapen-opptjening-pensjonsgivendeInntekt":              "pensjonsamhandling",
	"aapen-klager-klageOpprettet-DLT":                      "klage",
	"aapen-klager-klageOpprettet":                          "klage",
	"aapen-helse-sykepenger-utbetaling":                    "tbd",
	"aapen-grafana-paneler-v1":                             "tbd",
	"aapen-behandlingsgrunnlag-v2":                         "teamdatajegerne",
	"aapen-arena-dagpengeoppgaveLogget-v1-p":               "teamarenanais",
	"aapen-arena-14aVedtakIverksatt-v1-p":                  "teamarenanais",
	"aapen-arbeid-arbeidssoker-kontaktbruker-opprettet-q1": "paw",
	"aapen-arbeid-arbeidssoker-henvendelse-opprettet-p":    "paw",
	"StillingIntern":                                       "teampam",
	"StillingImportApiAdState":                             "teampam",
	"StillingEkstern":                                      "teampam",
	"Annonsehistorikk":                                     "teampam",
}
