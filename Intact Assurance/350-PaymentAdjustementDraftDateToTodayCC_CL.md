### 350-PaymentAdjustementDraftDateToTodayCC_CL

account : TEST_EN_US
          TEST_EN_US3
          TEST_EN_US
          TEST_EN_US
          
          
1000000196 pass
1000000198 pass
1000000200 pass
1000000202 pass
1000000206 pass
1000000208 pass
1000000210 pass
1000000212 pass

constat : quand le wizard d'account est "reset" soit il commence a 1000000150 ou lorsqu'il atteint 1000000150 l'account doit etre configurer differement et le test en question plante. cela semble etre la raison pourquoi le test fail au deux mois, il doit avoir le temps necessaire pour retomber a 0 et atteindre le compte 150 dans le wizard