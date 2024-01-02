## 🎯 Objectives
- Use logistic regression for binary classification
- Implement logistic regression for binary classification
- Address overfitting using regularization, to improve model performance

## Logistic Regression
![logistic-regression](.images/classification-1.png)  
Logistic regression uses the same basic formula as linear regression but it is regressing the probability of a categorical outcome. It gives a continuous value of `P(Y=1)` for a given `x`, which is converted to Y=0 or 1 based on a threshold value.

$$
    g_{w,b}(x)=\frac{1}{1+e^{f_{w,b}(x)}}
$$  

$$
    f_{w,b}(x)=wx+b
$$

### Terms
- **sigmoid function**: a function that turns linear regression to probability regression
- **threshold**: a value that decides P(Y=1) or P(Y=0)

### Decision Boundary
![decision-boundary](.images/classification-2.png)  
Decision boundary is a regression function when its output makes the sigmoid function to be threshold. If the threshold is 0.5, then the decision boundary should be $f_{w,b}(x)=0$, since $g(x)=\frac{1}{1+e^{-f(x)}}$.

## Gradient Descent
### Cost Function

$$
J(w,b)=\frac{1}{m}\sum^m{Loss(f_{w,b}(x^{(i)}), y^{(i)})}
$$

Cost function now one step further generalized by introducing **Loss** function, which measures difference of a single example to its target value.

$$
\begin{equation}
  Loss(f_{w,b}{(x)}, y)=\begin{cases}
    -\log{f}, & \text{if $y=1$}\\
    -\log{(1-f)}, & \text{if $y=0$}
  \end{cases}
\end{equation}
$$

## Questions
- How do you decide degree of polynomial function?