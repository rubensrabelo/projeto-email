## Guia de Configuração do Keycloak (Pós-Container)## 1. Realm Inicial

* Nome: provider

## 2. Cliente para Integração (Machine-to-Machine)

* ID do Cliente: emailn
* Protocolo: OpenID Connect
* Capability Config:
* Client Authentication: On (Habilita o uso de Client Secret).
   * Authorization: On.
   * Authentication Flow: Marque Service accounts roles (essencial para M2M).
* Credentials: Copie o Client Secret para usar nas suas variáveis de ambiente.

## 3. Configuração de Escopos e Mappers (Dedicados)

* Client Scopes (Geral): Alterar todos para Optional para evitar tokens muito pesados.
* Ajuste de Audiência (no emailn-dedicated):
* Configure a new mapper > Escolher Audience.
   * Name: aud_mapper.
   * Included Custom Audience: emailn.
   * Add to access token: On.
* Limpeza: Remova ou torne opcionais os atributos dedicados desnecessários no escopo.

## 4. Cliente para Usuários Finais

* ID do Cliente: emailn_users
* Configurações de Realm: Em Realm Settings > General, habilite a opção de usar Email as username (se desejar que o login seja o e-mail).
* User Profile: Em Realm Settings > User Profile, crie os atributos personalizados que você deseja que existam no sistema.

## 5. Gestão de Usuários

* Criação: Criar o usuário e definir a senha em Credentials.
* Senha: Desmarcar a opção Temporary (para não exigir troca no primeiro login via Insomnia).
* Atributos: Na aba Attributes do usuário, preencha os valores criados no passo anterior.
* Exposição no Token: Dentro de Clients > emailn_users > Client Scopes > dedicated, use o Configure a new mapper do tipo User Attribute para mapear o atributo do usuário para o JWT.

## 6. Ajustes de Sessão (Tokens)

* Em Realm Settings > Tokens, ajuste o Access Token Lifespan conforme sua necessidade de teste (ex: 1 hora) para evitar que o token expire muito rápido durante o desenvolvimento.

