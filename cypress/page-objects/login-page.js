import selectors from "../static/selectors";

const rancherURL = Cypress.config("baseUrl");

export function login(userName, password) {
  cy.visit(rancherURL);
  cy.get(selectors.login.username).type(userName);
  cy.get(selectors.login.password).type(password);
  cy.get(selectors.login.submit).click({ force: true });
  cy.url().should('include', '/dashboard');
}
