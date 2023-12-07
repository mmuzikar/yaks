package org.citrusframework.yaks.camelk.suggestions;

import org.citrusframework.yaks.camelk.CamelKSettings;
import org.citrusframework.yaks.camelk.CamelKSupport;
import org.citrusframework.yaks.camelk.VariableNames;
import org.citrusframework.yaks.camelk.model.Integration;
import org.citrusframework.yaks.camelk.model.IntegrationList;
import org.citrusframework.yaks.kubernetes.KubernetesResource;
import org.citrusframework.yaks.kubernetes.KubernetesSettings;
import org.citrusframework.yaks.kubernetes.KubernetesSupport;

import java.util.List;
import java.util.stream.Collectors;

import org.citrusframework.Citrus;
import org.citrusframework.CitrusInstanceManager;

import io.fabric8.kubernetes.client.KubernetesClient;
import io.github.mmuzikar.interactive.cucumber.api.SuggestionItem;
import io.github.mmuzikar.interactive.cucumber.api.SuggestionProvider;

public class IntegrationNameSuggestionProvider implements SuggestionProvider {

    KubernetesClient client;
    Citrus citrus;

    public IntegrationNameSuggestionProvider() {
        citrus = CitrusInstanceManager.getOrDefault();
        client = KubernetesSupport.getKubernetesClient(citrus);
    }

    @Override
    public List<SuggestionItem> provide(String step) {
        return client.resources(Integration.class)
            .inNamespace(getNamespace())
            .list()
            .getItems()
            .stream().map(it -> SuggestionItem.withText(it.getMetadata().getName())).collect(Collectors.toList());
    }

    String getNamespace() {
        return (String) citrus.getCitrusContext().getGlobalVariables().getVariables()
            .getOrDefault(VariableNames.CAMELK_NAMESPACE, KubernetesSettings.getNamespace());
    }
}
