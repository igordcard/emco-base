// SPDX-License-Identifier: Apache-2.0
// Based on Code: https://github.com/johandry/klient

package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/resource"
)

const (
	// Period of time in seconds given to the resource to terminate gracefully when delete it (used when require to recreate the resource). Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when force is true (force deletion)
	gracePeriod = -1
	// If true, cascade the deletion of the resources managed by this resource (e.g. Pods created by a ReplicationController).
	cascade = true
)

// Delete creates a resource with the given content
func (c *Client) Delete(content []byte) error {
	r := c.ResultForContent(content, nil)
	return c.DeleteResource(r)
}

// DeleteFiles create the resource(s) from the given filenames (file, directory or STDIN) or HTTP URLs
func (c *Client) DeleteFiles(filenames ...string) error {
	r := c.ResultForFilenameParam(filenames, nil)
	return c.DeleteResource(r)
}

// DeleteResource applies the given resource. Create the resources with `ResultForFilenameParam` or `ResultForContent`
func (c *Client) DeleteResource(r *resource.Result) error {
	if err := r.Err(); err != nil {
		return err
	}
	return r.Visit(deleteResource)
}

func deleteResource(info *resource.Info, err error) error {
	if err != nil {
		return failedTo("deleteResource", info, err)
	}

	// TODO: Background or Foreground?
	// policy := metav1.DeletePropagationForeground
	policy := metav1.DeletePropagationBackground
	options := metav1.DeleteOptions{
		PropagationPolicy: &policy,
	}

	if _, err := deleteWithOptions(info, &options); err != nil {
		return failedTo("delete", info, err)
	}
	return nil
}

func defaultDeleteOptions() *metav1.DeleteOptions {
	// TODO: Change DryRun value when DryRun is implemented
	dryRun := false

	options := &metav1.DeleteOptions{}
	if gracePeriod >= 0 {
		options = metav1.NewDeleteOptions(int64(gracePeriod))
	}

	if dryRun {
		options.DryRun = []string{metav1.DryRunAll}
	}

	// TODO: Background or Foreground?
	// policy := metav1.DeletePropagationBackground
	policy := metav1.DeletePropagationForeground
	if !cascade {
		policy = metav1.DeletePropagationOrphan
	}
	options.PropagationPolicy = &policy

	return options
}

func deleteWithOptions(info *resource.Info, options *metav1.DeleteOptions) (runtime.Object, error) {
	if options == nil {
		options = defaultDeleteOptions()
	}
	return resource.NewHelper(info.Client, info.Mapping).DeleteWithOptions(info.Namespace, info.Name, options)
}
