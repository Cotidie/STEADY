## 🎯 Objectives
- Define supervised/unsupervised machine learning
- Define a regression model
- Implement/Optimize gradient descent

## Overview
> Field of study that gives computers the ability to learn without being explicitly programmed.
>  
> by Arthor Samuel

## Supervised/Unsupervised
### Supervised
![Alt text](./.images/introduction-supervised-1.png)  
Supervised learning refers to algorithms that map input to output or `x` to `y` by learning from the given "right answer" dataset. It's once more divided into `Regression` and `Classification`.
- **Regression**: Predicts a number from the approximation function
- **Classification**: Predicts a class from multiple inputs using boundary function
  
### Unsupervised
![unsupervised learning](./.images/introduction-unsupervised-1.png)  
Unsupervised learning refers to algorithms that take only a set of inputs and find structure in the given data. It doesn't requre right answers or labels.
- **Algorithms**: Clustering, Anomaly Detection, Dimension Reduction
- ex) news grouping, DNA Microarray...

## Linear Regression
### Term
- **Training Set**: Data used to train the model
  - feature/target: input `x`, output `y`
  - $(x^{(i)}, y^{(i)})$: `i`th example
- **Model**: Regression function
  - $f_{w,b}(x)$: The model $f(x)=wx+b$
  - $\hat{y}$: an estimate from a model
  - `w`, `b`: parameters, coefficients

### Cost Function
Cost function $J(w,b)$is a function that measures how accurate the model is compared to the training set. Below is `Squred Error Cost Function` which is the most commonly used one for linear regression.

$$
J(w,b)=\frac{1}{2m}\sum_{i}^{m}{(\hat{y}^{i}-y^{(i)})^2}
$$

The goal is to **minimize $J(w,b)$** using varius methods, mainly using `Gradient Descent`.  
![cost-function](.images/introduction-cost-function-1.png)  

## Gradient Descent
