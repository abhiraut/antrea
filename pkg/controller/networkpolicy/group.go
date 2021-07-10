// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkpolicy

import (
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"

	crdv1alpha1 "antrea.io/antrea/pkg/apis/crd/v1alpha1"
)

// addGroup is responsible for processing the ADD event of a Group resource.
func (n *NetworkPolicyController) addGroup(curObj interface{}) {
	g := curObj.(*crdv1alpha1.Group)
	klog.V(2).Infof("Processing ADD event for Group %s", g.Name)
}

// updateGroup is responsible for processing the UPDATE event of a Group resource.
func (n *NetworkPolicyController) updateGroup(oldObj, curObj interface{}) {
	cg := curObj.(*crdv1alpha1.Group)
	klog.V(2).Infof("Processing UPDATE event for Group %s", cg.Name)
}

// deleteGroup is responsible for processing the DELETE event of a Group resource.
func (n *NetworkPolicyController) deleteGroup(oldObj interface{}) {
	og, ok := oldObj.(*crdv1alpha1.Group)
	klog.V(2).Infof("Processing DELETE event for Group %s", og.Name)
	if !ok {
		tombstone, ok := oldObj.(cache.DeletedFinalStateUnknown)
		if !ok {
			klog.Errorf("Error decoding object when deleting Group, invalid type: %v", oldObj)
			return
		}
		og, ok = tombstone.Obj.(*crdv1alpha1.Group)
		if !ok {
			klog.Errorf("Error decoding object tombstone when deleting Group, invalid type: %v", tombstone.Obj)
			return
		}
	}
}
