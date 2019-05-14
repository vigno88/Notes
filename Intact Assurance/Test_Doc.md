# Test Learning

## General Info

 - [Jenkins Url](https://stha38e56:444/view/Billing/)

- The *schedule test* flow is used to shedule and rul all automated tests
  - Once in the flow, we can see a list of build
    - BUILD_ENDTOEND_GWBC-BDI = gobrio (old system that we want to replace by Guidewire)
    - BUILD_ENDTOEND_GWBC = Guidewire
    - D1 <- D2 <- D3
- To see the result of the test int the ITA gui, click on **Results URL**
- Do not note in the suite follow up anything related to **Nouvelle Affaire**
- Doc [ITA](https://apps.iad.ca.inet/sites/adsth/ITA/default.aspx )
## Tasks to do in the morning
[3.Scheduled Tests [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/)

Always verify if the fail already happened in the past
BDI in priority since other test depends of them later in the day

- [x] [BUILD_ENDTOEND_GWBC-BDI_D1_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D1_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_D2_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D2_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_D3_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D3_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_FLOW/)
- [x] Look at the ping in the most recent build in [TEST_Integration [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_Integration/)
> 12 fail in TestNonPROD
    
- [x] [TEST_SERVICES_JMeter [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_SERVICES_JMeter/)
> BIA (31) and BSA (1)
- [ ] [TESTS_ALL_GWBC_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_ALL_GWBC_FLOW/)
> **8 fails**
>  - 1276-GenerateOutgoingChequeRefunds_Flow_D1
>  - 1276-GenerateOutgoingChequeRefunds_Flow_D2
>  - **558-SalesAgentRegisteringCC**
>  - 1276-GenerateOutgoingChequeRefunds_Flow_D3
>  -  83742-CommissionPayment_CL 
>  - Test langue adler
- [x] [TESTS_billing-BP-Clock [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-BP-Clock/)
- [x] [TESTS_billing-delinquency_part1 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part1/)
- [x] [TESTS_billing-delinquency_part2 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part2/)
- [x] [TESTS_Billing-statement [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_Billing-statement/)
> Still running
- [x] [TESTS_WORKFLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_WORKFLOW/)
## How to fill Testsuit follow-up
- Test Suite
  - Name of the suite in Jenkins
- Test failing
  - Can be found in the HTML report
- Branch
  - Can be found in the HTML report
- Date
  - suite dashboard
- Environment
  - **AppConfigEnv** in the dashboard of the test report
- Report
  - The url of the test report
- Execution host
  - **Hostname** in the test report
- Screen def
  - In the test report, second menu from the top, in the test scrolldown menu, find the test that failed. In the description of the test, the screendef can be found in the step before the one that failed.
- Test fails since:
  - Date of the failure
- Detail:
  - Conclusion of the investigation