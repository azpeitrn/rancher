import {login} from "../../page-objects/login-page";
import selectors from "../../static/selectors";


describe('Access Rancher home web page', () => {

    beforeEach(() => {
        login('rancher-standard', 'WaczB8tSX0BppPyyYAtU' );
    });

    it('Checks if the main web page opens up.', () => {
        cy.get(selectors.home.main).should('be.visible');
    });

    it('Checks if the main web page title is correct', () => {
        cy.title().should('eq', 'Rancher');
    });
});
