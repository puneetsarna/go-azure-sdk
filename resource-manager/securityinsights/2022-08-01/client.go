package v2022_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/actions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/alertrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/alertruletemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/automationrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/bookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/dataconnectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/entitytypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/incidentalerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/incidentbookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/incidentcomments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/incidententities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/incidentrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/incidents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/sentinelonboardingstates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/threatintelligence"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/watchlistitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2022-08-01/watchlists"
)

type Client struct {
	Actions                  *actions.ActionsClient
	AlertRuleTemplates       *alertruletemplates.AlertRuleTemplatesClient
	AlertRules               *alertrules.AlertRulesClient
	AutomationRules          *automationrules.AutomationRulesClient
	Bookmarks                *bookmarks.BookmarksClient
	DataConnectors           *dataconnectors.DataConnectorsClient
	EntityTypes              *entitytypes.EntityTypesClient
	IncidentAlerts           *incidentalerts.IncidentAlertsClient
	IncidentBookmarks        *incidentbookmarks.IncidentBookmarksClient
	IncidentComments         *incidentcomments.IncidentCommentsClient
	IncidentEntities         *incidententities.IncidentEntitiesClient
	IncidentRelations        *incidentrelations.IncidentRelationsClient
	Incidents                *incidents.IncidentsClient
	SentinelOnboardingStates *sentinelonboardingstates.SentinelOnboardingStatesClient
	ThreatIntelligence       *threatintelligence.ThreatIntelligenceClient
	WatchlistItems           *watchlistitems.WatchlistItemsClient
	Watchlists               *watchlists.WatchlistsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	actionsClient := actions.NewActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&actionsClient.Client)

	alertRuleTemplatesClient := alertruletemplates.NewAlertRuleTemplatesClientWithBaseURI(endpoint)
	configureAuthFunc(&alertRuleTemplatesClient.Client)

	alertRulesClient := alertrules.NewAlertRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&alertRulesClient.Client)

	automationRulesClient := automationrules.NewAutomationRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&automationRulesClient.Client)

	bookmarksClient := bookmarks.NewBookmarksClientWithBaseURI(endpoint)
	configureAuthFunc(&bookmarksClient.Client)

	dataConnectorsClient := dataconnectors.NewDataConnectorsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataConnectorsClient.Client)

	entityTypesClient := entitytypes.NewEntityTypesClientWithBaseURI(endpoint)
	configureAuthFunc(&entityTypesClient.Client)

	incidentAlertsClient := incidentalerts.NewIncidentAlertsClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentAlertsClient.Client)

	incidentBookmarksClient := incidentbookmarks.NewIncidentBookmarksClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentBookmarksClient.Client)

	incidentCommentsClient := incidentcomments.NewIncidentCommentsClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentCommentsClient.Client)

	incidentEntitiesClient := incidententities.NewIncidentEntitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentEntitiesClient.Client)

	incidentRelationsClient := incidentrelations.NewIncidentRelationsClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentRelationsClient.Client)

	incidentsClient := incidents.NewIncidentsClientWithBaseURI(endpoint)
	configureAuthFunc(&incidentsClient.Client)

	sentinelOnboardingStatesClient := sentinelonboardingstates.NewSentinelOnboardingStatesClientWithBaseURI(endpoint)
	configureAuthFunc(&sentinelOnboardingStatesClient.Client)

	threatIntelligenceClient := threatintelligence.NewThreatIntelligenceClientWithBaseURI(endpoint)
	configureAuthFunc(&threatIntelligenceClient.Client)

	watchlistItemsClient := watchlistitems.NewWatchlistItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&watchlistItemsClient.Client)

	watchlistsClient := watchlists.NewWatchlistsClientWithBaseURI(endpoint)
	configureAuthFunc(&watchlistsClient.Client)

	return Client{
		Actions:                  &actionsClient,
		AlertRuleTemplates:       &alertRuleTemplatesClient,
		AlertRules:               &alertRulesClient,
		AutomationRules:          &automationRulesClient,
		Bookmarks:                &bookmarksClient,
		DataConnectors:           &dataConnectorsClient,
		EntityTypes:              &entityTypesClient,
		IncidentAlerts:           &incidentAlertsClient,
		IncidentBookmarks:        &incidentBookmarksClient,
		IncidentComments:         &incidentCommentsClient,
		IncidentEntities:         &incidentEntitiesClient,
		IncidentRelations:        &incidentRelationsClient,
		Incidents:                &incidentsClient,
		SentinelOnboardingStates: &sentinelOnboardingStatesClient,
		ThreatIntelligence:       &threatIntelligenceClient,
		WatchlistItems:           &watchlistItemsClient,
		Watchlists:               &watchlistsClient,
	}
}
