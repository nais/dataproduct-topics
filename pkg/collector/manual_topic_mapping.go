package collector

var manualTopicMapping = map[string]string{
	"aapen-pdl-adresse-v1":                                   "pdl",
	"dittnav-brukernotifikasjon-done-v1-backup":              "personbruker",
	"nais-deployment-requests":                               "aura",
	"nais-deployment-status":                                 "aura",
	"privat-tiltaksgjennomforing-godkjentAvtale":             "arbeidsgiver",
	"privat-tiltaksgjennomforing-statistikkformidling":       "arbeidsgiver",
	"privat-tiltaksgjennomforing-smsVarselResultat":          "arbeidsgiver",
	"privat-tiltaksgjennomforing-smsVarsel":                  "arbeidsgiver",
	"privat-tortuga-sekvensnummerTilstand":                   "pensjonsamhandling",
	"privat-tortuga-skatteoppgjorhendelse":                   "pensjonsamhandling",
	"privat-sykepengebehandling":                             "tbd",
	"privat-syfooversikt-tilfelle-retry-v1":                  "teamsykefravr",
	"privat-sf-person-v1":                                    "crm", // Team CRM?, no nais team slug,
	"privat-sf-person-pilot-v1":                              "crm", // Team CRM?, no nais team slug,
	"nks-sf-pdl-v1":                                          "crm", // Team CRM?, no nais team slug,
	"privat-person-pdl-aktor-v1":                             "pdl",
	"privat-person-tpsHendelse-mottatt-v1":                   "pdl",
	"privat-person-pdl-lineage-v1":                           "pdl",
	"privat-person-pdl-identhendelseSletting-v1":             "pdl",
	"privat-person-pdl-hendelsesaggregat-v2":                 "pdl",
	"privat-person-pdl-fregTilbakemelding-v1":                "pdl",
	"privat-person-pdl-e500-linje-v1":                        "pdl",
	"person-privat-pdl-e500-linje-v1":                        "pdl",
	"privat-person-navMottakStatus-v1":                       "pdl",
	"privat-person-mottakEndring-v1":                         "pdl",
	"privat-person-fregTilbakemelding-mottatt-v1":            "pdl",
	"privat-pdl-sfobjects-v1":                                "pdl",
	"privat-k9-dittnav-varsel-beskjed":                       "pleiepenger",
	"privat-eessipensjon-krav-initialisering":                "eessipensjon",
	"privat-dagpenger-journalpost-mottatt-alpha":             "teamdagpenger",
	"privat-dagpenger-behov-alpha":                           "teamdagpenger",
	"postnummer":                                             "tbd",
	"pensjon-pdl-kafka-retry":                                "pensjon-saksbehandling",
	"nada-connect-status":                                    "aura",
	"nada-connect-offsets":                                   "aura",
	"nada-connect-configs":                                   "aura",
	"kafka-monitor":                                          "aura",
	"fpsoknad-mottak":                                        "teamforeldrepenger",
	"fpinfo-historikk":                                       "teamforeldrepenger",
	"dittnav-brukernotifikasjon-nyOppgave-v1-backup":         "personbruker",
	"dittnav-brukernotifikasjon-nyBeskjed-v1-backup":         "personbruker",
	"dittnav-brukernotifikasjon-cached-done-v1-backup":       "personbruker",
	"aapen-opptjening-pensjonsgivendeInntekt":                "pensjonsamhandling",
	"aapen-klager-klageOpprettet-DLT":                        "klage",
	"aapen-klager-klageOpprettet":                            "klage",
	"aapen-helse-sykepenger-utbetaling":                      "tbd",
	"aapen-grafana-paneler-v1":                               "tbd",
	"aapen-behandlingsgrunnlag-v2":                           "teamdatajegerne",
	"aapen-arena-dagpengeoppgaveLogget-v1-p":                 "teamarenanais",
	"aapen-arena-14aVedtakIverksatt-v1-p":                    "teamarenanais",
	"aapen-arbeid-arbeidssoker-kontaktbruker-opprettet-q1":   "paw",
	"aapen-arbeid-arbeidssoker-henvendelse-opprettet-p":      "paw",
	"StillingIntern":                                         "teampam",
	"StillingImportApiAdState":                               "teampam",
	"StillingEkstern":                                        "teampam",
	"Annonsehistorikk":                                       "teampam",
	"aapen-registre-medlemskapEndret-v1-p":                   "team-rocket",
	"public-ereg-cache-org-json":                             "arbeidsforhold",
	"privat-pip-idm-ansatt":                                  "idm",
	"privat-pip-pdl-aktorid-q0":                              "pdl",
	"privat-pip-pdl-aktorid":                                 "pdl",
	"privat-pip-pdl-aktorid-q1":                              "pdl",
	"privat-pip-pdl-aktorid-q2":                              "pdl",
	"privat-pip-pdl-fnr":                                     "pdl",
	"privat-pip-pdl-fnr-q0":                                  "pdl",
	"privat-pip-pdl-fnr-q1":                                  "pdl",
	"privat-pip-pdl-fnr-q2":                                  "pdl",
	"privat-pip-pdl-gt-q0":                                   "pdl",
	"privat-pip-pdl-gt-q1":                                   "pdl",
	"privat-pip-pdl-gt-q2":                                   "pdl",
	"privat-pip-pdl-gt-v2-q0":                                "pdl",
	"privat-pip-pdl-gt-v2-q1":                                "pdl",
	"privat-pip-pdl-gt-v2-q2":                                "pdl",
	"privat-pip-pdl-person":                                  "pdl",
	"privat-pip-pdl-person-q0":                               "pdl",
	"privat-pip-pdl-person-q1":                               "pdl",
	"privat-pip-pdl-person-q2":                               "pdl",
	"aapen-arena-14aVedtakIverksatt-v1-q1":                   "teamarenanais",
	"aapen-arena-dagpengeoppgaveLogget-v1-q1":                "teamarenanais",
	"aapen-oppfolging-vedtakSendt-v1-q0":                     "pto",
	"aapen-pdl-geografisktilknytning-v1":                     "pdl",
	"privat-pip-stream-pdl-aktorid":                          "pdl",
	"privat-stream-pip-pdl-aktorid":                          "pdl",
	"privat-stream-pip-pdl-fnr":                              "pdl",
	"privat-stream-pip-pdl-gt":                               "pdl",
	"privat-stream-pip-pdl-gt-v2":                            "pdl",
	"privat-stream-pip-pdl-person":                           "pdl",
	"privat-pgi-hendelse":                                    "pensjonopptjening",
	"privat-eessipensjon-krav-initialisering-q2":             "eessipensjon",
	"kafka-faggruppe-apabepa":                                "team-soknad",
	"dvh-sykm-dlq":                                           "disykefravar",
	"aapen-arbeid-arbeidssoker-kontaktbruker-opprettet-q0":   "paw",
	"privat-pgi-inntekt":                                     "pensjonopptjening",
	"privat-joark-dagpengesoknaderEksportert-v1-t1":          "teamdagpenger",
	"privat-dagpenger-subsumsjon-brukt-q2":                   "teamdagpenger",
	"aapen-registre-medlemskapEndret-v1-q2":                  "team-rocket",
	"privat-jfr-manuell-oppretter-feilUnderBehandling-alpha": "isa",
	"privat-dagpenger-subsumsjon-brukt-data-q2":              "teamdagpenger",
	"eessipensjon-statistikk-q2":                             "eessipensjon",
	"privat-dele-omsorgsdager-melding-preprosessert":         "pleiepenger",
	"privat-dagpenger-subsumsjon-brukt":                      "teamdagpenger",
	"privat-person-tpsHendelse-mottatt-v1-q2":                "pdl",
	"privat-person-tpsHendelse-mottatt-v1-q1":                "pdl",
	"privat-pgi-nextsekvensnummer":                           "pensjonopptjening",
	"privat-dagpenger-subsumsjon-brukt-q0":                   "teamdagpenger",
	"privat-isa-jfrBehandlingsfeil-q1":                       "isa",
	"privat-dagpenger-subsumsjon-brukt-data-q0":              "teamdagpenger",
	"privat-sif-innsyn-mottak":                               "pleiepenger",
}
